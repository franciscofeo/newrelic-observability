# Monitoring and Observability with Prometheus and New Relic

The application will run in `localhost:8080`. 

Grafana username and password is `admin`.

### Grafana Screenshots

![Captura de tela 2025-02-19 201512](https://github.com/user-attachments/assets/82b496ca-cbd2-4628-93be-75a5c17ed320)

### Prometheus Screenshots

![Captura de tela 2025-02-19 105025](https://github.com/user-attachments/assets/b8517145-42ad-440a-a1fe-689529e6e89d)

![prometheus](https://github.com/user-attachments/assets/d5c238b5-893a-4625-af3c-b1ef99fdf645)


### New Relic Screenshots

![image](https://github.com/user-attachments/assets/944df70c-10f7-4c55-bcd1-f4e3f7058cf7)

![image](https://github.com/user-attachments/assets/510c1bc1-ec98-440b-ad8c-9fa0319b3360)

![image](https://github.com/user-attachments/assets/5898c62c-00a8-4708-a68a-fc1a2452b982)

### Application endpoints

Create a record:
```bash
curl --location 'localhost:8080/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "John Doe"
}'
```

Return all the records:
```bash
curl --location 'localhost:8080/records'
```

Delete a record:
```bash
curl --location --request DELETE 'localhost:8080/delete?id=id'
```
