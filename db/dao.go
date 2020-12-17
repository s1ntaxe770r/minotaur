package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Connect : Instantiate db connection
func Connect() *sql.DB {
	user := os.Getenv("PG_USER")
	// pass := os.Getenv("PG_PASS")
	dbname := os.Getenv("DB_NAME")
	constr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", user, dbname)
	db, err := sql.Open("postgres", constr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
