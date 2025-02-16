package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

var db *sql.DB

type Record struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func initDB() {
	log.Println("Initializing PostgreSQL Database.")
	var err error
	connStr := "user=postgres host=postgres dbname=cruddb password=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS records (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func addRecord(w http.ResponseWriter, r *http.Request) {
	var record Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Creating record: " + record.Name)
	query := `INSERT INTO records (name) VALUES ($1) RETURNING id, created_at`
	err := db.QueryRow(query, record.Name).Scan(&record.ID, &record.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getAllRecords(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, created_at FROM records")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Println("Retrieving all records")
	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Name, &record.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	log.Println("Deleting record: " + id)
	_, err := db.Exec("DELETE FROM records WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("observability-app"),
		newrelic.ConfigLicense("___"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		fmt.Println(err)
	}
	writer := logWriter.New(os.Stdout, app)
	logger := log.New(&writer, "", log.Default().Flags())

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/add", addRecord))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/records", getAllRecords))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/delete", deleteRecord))

	logger.Println("Server is running on port 8080. Using the New Relic Logger Forward.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
