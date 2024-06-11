package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var version string
	if err := db.QueryRow("SELECT version()").Scan(&version); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connection in version=%s\n", version)
	return db
}
