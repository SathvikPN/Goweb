# Goweb
Get started with web application development using Go

### setup
- Router: github.com/gorilla/mux
- Database: Postgres
- Database driver for Go: github.com/lib/pq 
- Environment Variable Loader: github.com/joho/Godotenv

### database instance
- postgres database as docker container. (docker-compose up -d)     [--env-file [.env]]
- create database tables
- if database container hosted in cloud, use credentials in .env file (create file). sample .env file is .env.example
