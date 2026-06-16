## Day 3 - PostgreSQL

# Theory 
# Why not write directly in Go? 
- Bad: databaseURL := "postgres://postgres:password@localhost:5432/wallet"
- Good: databaseURL := os.Getenv("DATABASE_URL")
- The Go code changes when database address, password, or environment. The os.Getenv reads and environment variable and return empty string when its absent

# Steps I took
1. Created DATABASE_URL
2. Installed pgx
3. Read DATABASE_URL in main.go
4. Create connection pool 
5. Made some changes in main.go to set DATABASE_URL

# The code that prepeares go to connect with postgre
1. Creates a context
2. Creates a reusable pool of database connections 
3. Stops program is pool setup is failed
4. Ensures pool is closed when surrounding functions finishes

# Other learned things
- Context is a standard golang library that is responsible for: 
    1. Cancellation: "Stop this operation" 
    2. Deadline: "This operation must finish before this time" 
    3. Timeout: "Stop after five seconds" and etc
