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
	pass := os.Getenv("PG_PASS")
	host := "localhost"
	dbname := os.Getenv("DB_NAME")
	port := "5432"
	constr := fmt.Sprintf("user=%s dbname=%s pass=%s port=%s host=%s sslmode=verify-full", user, dbname, pass, port, host)
	db, err := sql.Open("postgres", constr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Insert add a project to the db
func Insert() error {

}
