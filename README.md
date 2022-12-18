# XM-Golang Challenge
## Problem definition
Build a microservice to handle companies. It should provide CRUDING operation.
### Requirements
![Requirements](./diagram/XM Golang Exercise - v22.0.0 .pdf)

### System tests
1. Signup Endpoint: 
```bash
   curl --location --request POST 'localhost:8080/user/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"mahmoud@mekki.com",
    "password":"24072017",
    "user_name":"Mekki"
}'
```
2. Login Endpoint:
```bash
 curl --location --request GET 'localhost:8080/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"mahmoud@mekki.com",
    "password":"24072017",
}'
```
3. Create Company endpoint
```bash
  curl --location --request POST 'localhost:8080/company' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NzE0NzI5ODYsImlhdCI6MTY3MTMwMDE4NiwiaXNzIjoiTWVra2kifQ.N-6EK74bYPG-s0NpL06jRxsKmF-I-8b7SmXKcq_3bsBzFw3dbY3CtvIgo-9e2fhFIWQ-N5ALaHv2UPQgUoXAvw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mekki",
    "type":"NonProfit",
    "employees_number":1000,
    "registered": true
}'
```
4. Get Company by ID
```bash
curl --location --request GET 'localhost:8080/company/5f3a5f08-53a6-4ee8-8327-61f3fb6f78fb' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NzE0NzI5ODYsImlhdCI6MTY3MTMwMDE4NiwiaXNzIjoiTWVra2kifQ.N-6EK74bYPG-s0NpL06jRxsKmF-I-8b7SmXKcq_3bsBzFw3dbY3CtvIgo-9e2fhFIWQ-N5ALaHv2UPQgUoXAvw' \
--header 'Content-Type: application/json'
```
5. Update Company Endpoint:
```bash
curl --location --request PATCH 'localhost:8080/company/5f3a5f08-53a6-4ee8-8327-61f3fb6f78fb' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NzE0NzI5ODYsImlhdCI6MTY3MTMwMDE4NiwiaXNzIjoiTWVra2kifQ.N-6EK74bYPG-s0NpL06jRxsKmF-I-8b7SmXKcq_3bsBzFw3dbY3CtvIgo-9e2fhFIWQ-N5ALaHv2UPQgUoXAvw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "employees_number":100
}'
```
6. Delete Company Endpoint 
```bash
curl --location --request DELETE 'localhost:8080/company/5f3a5f08-53a6-4ee8-8327-61f3fb6f78fb' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NzE0NzI5ODYsImlhdCI6MTY3MTMwMDE4NiwiaXNzIjoiTWVra2kifQ.N-6EK74bYPG-s0NpL06jRxsKmF-I-8b7SmXKcq_3bsBzFw3dbY3CtvIgo-9e2fhFIWQ-N5ALaHv2UPQgUoXAvw' \
--header 'Content-Type: application/json'
```

## To Run the Tests Locally
# Golang SDK <= 1.17
```bash
> go mod tidy
> go mod vendor
> docker build ./
> docker-compose up [runinng mysql testing server]
> go test tests/companies/* [tests]
```
# Golang SDK >= 1.17
```bash
> docker-compose up [runinng mysql testing server]
> go test tests/companies/* [tests]
```

## Run the system on docker
```bash
> docker-compose up 
```