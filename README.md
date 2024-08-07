# Goweb
A simple backend development project using Go

> Thunder Client configuration available at tool/ directory

### Run
- configure .env file using .env.example structure, source .env
- docker-compose up -d database 
- docker exec -it db-container-id psql -U a_user -d goweb
- `go run .`

### setup
- Router: github.com/gorilla/mux
- Database: Postgres
- Database driver for Go: github.com/lib/pq 
- Environment Variable Loader: github.com/joho/Godotenv

### database instance
- postgres database as docker container. (docker-compose up -d service_name)
- create database tables
- if database container hosted in cloud, use credentials in .env file (create file). sample .env file is .env.example

### models
- package to store database schema
- DB schema is mapped as golang structs (exported, json tag included)

### middleware
- bridge between APIs and Database
- middleware are handlers to handle DB operations (CRUD)
- add CreatePost handler and logic function to persist post in database

### router
- includes mapping of REST API endpoints to handlers

### entrypoint
- invoke server and expose router endpoints
