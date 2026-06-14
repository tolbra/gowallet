# Day 1 - Current project state

## Works

- Go server: works  
- Current port: 3030
- Working routes: GET http://localhost:3030/healthz
- Docker: works; Compose successfully started the services 
- PostgreSQL: works; container status is running 

## Fails

- Error: unable to get image postgres:17 
- Command that caused it: docker compose up -d
- Cause: Docker Desktop was not running 
- Fix: Start Docker Desktop and run the command again

## Current structure 

- Server starts from: ./cmd/api/main.go
- Router package: github.com/go-chi/chi/v5
- Logging middleware: github.com/go/chi/chi/v5/middleware
- Database Docker image: postgres:17
- PostgreSQL user: postgres
- PostgreSQL password: password
- PostgreSQL database: wallet
- PostgreSQL port: 5432
- Go database connection: not implemented yet
- Future database URL: postgres://postgres:password@localhost:5432/wallet?sslmode=disable  