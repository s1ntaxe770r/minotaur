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
	var db *sql.DB
	defer db.Close()
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
func Insert(project *Project, dbcon *sql.DB) error {
	sqlstmnt := `insert into "projects"("name", "liveurl","github") values($1, $2,$3)`
	_, inserterr := dbcon.Exec(sqlstmnt, project.Name, project.LiveURL, project.Github)
	if inserterr != nil {
		return inserterr
	}
	return nil
}

// QueryAll  returns all projects
func QueryAll(dbcon *sql.DB, projects *Projects) error {
	rows, err := dbcon.Query(`
		SELECT
			id,
			name,
			live_url,
			github
		FROM projects`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		proj := Project{}
		err = rows.Scan(
			&proj.ID,
			&proj.Name,
			&proj.LiveURL,
			&proj.Github,
		)
		if err != nil {
			return err
		}
		projects.projects = append(projects.projects, proj)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// QueryOne retrieve a single project
func QueryOne(dbcon *sql.DB, project *Project, id string) error {
	sqltmnt := `SELECT id ,name,live_url,github FROM projects WHERE id=$1 ORDER BY id DESC LIMIT 1`
	row := dbcon.QueryRow(sqltmnt, id)
	err := row.Scan(
		project.ID,
		project.Name,
		project.LiveURL,
		project.Github,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete  removes project from db
func Delete(dbcon *sql.DB, id string) error {
	sqlstmnt := `delete from "projects" where id=$1`
	_, err := dbcon.Exec(sqlstmnt, id)
	if err != nil {
		return nil
	}
	return nil
}

// Update update a project
func Update(dbcon *sql.DB, project *Project, id string) (*Project, error) {
	sqltmnt := `UPDATE projects SET  name=$1, live_url=$2 , github=$3 WHERE id=$4`
	_, err := dbcon.Exec(sqltmnt, project.Name, project.LiveURL, project.Github, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}
