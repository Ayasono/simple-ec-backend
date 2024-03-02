package database

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (
  dbDriver = "postgres"
  dbSource = "postgresql://root:123456@localhost:5432/kins_db?sslmode=disable"
)

// ConnectDB Connect initializes and returns a database connection
func ConnectDB() *sql.DB {
  db, err := sql.Open(dbDriver, dbSource)
  if err != nil {
    log.Fatal("Error connecting to database:", err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal("Error pinging database:", err)
  }

  fmt.Println("Successfully connected to database")
  return db
}
