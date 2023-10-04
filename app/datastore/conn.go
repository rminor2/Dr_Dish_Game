package gamedb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dbInst *sql.DB

func init() {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	dbInst = db
}

func RWInstance() *sql.DB {
	return dbInst
}
