package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/s1ntaxe770r/PPI/utils"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	dblogger = utils.DBlogger()
)

//Connect : Instantiate db connection
func Connect() *sql.DB {
	// user := os.Getenv("PG_USER")
	// pass := os.Getenv("PG_PASS")
	// host := os.Getenv("DB_HOST")
	// dbname := os.Getenv("DB_NAME")
	// port := "5432"
	// constr := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	connstr := os.Getenv("CONNSTR")
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
		dblogger.Println(err)

	}
	return db
}

// Insert add a project to the db
func Insert(project Project, dbcon *sql.DB) error {
	sqlstmnt := `INSERT INTO projects(name,live_url,github) VALUES($1,$2,$3)`
	name := project.Name
	url := project.LiveURL
	github := project.Github
	dblogger.Println(name)
	res, inserterr := dbcon.Exec(sqlstmnt, name, url, github)
	dblogger.Println(res)
	if inserterr != nil {
		dblogger.Println(inserterr)
		return inserterr

	}
	return nil
}

// QueryAll  returns all projects
func QueryAll(dbcon *sql.DB, projects *Projects) error {
	rows, err := dbcon.Query(`SELECT * FROM projects`)
	if err != nil {
		dblogger.Println(err)
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
			dblogger.Println(err)
			return err
		}
		projects.projects = append(projects.projects, proj)
		dblogger.Println(projects.projects)
	}
	err = rows.Err()
	if err != nil {
		dblogger.Println(err)
		return err
	}
	return nil
}

// QueryOne retrieve a single project
func QueryOne(dbcon *sql.DB, project *Project, id string) error {
	sqltmnt := `SELECT id ,name,live_url,github FROM projects WHERE id=$1`
	row := dbcon.QueryRow(sqltmnt, id)
	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.LiveURL,
		&project.Github,
	)
	if err != nil {
		dblogger.Println(err)
		return err
	}
	return nil
}

// Delete  removes project from db
func Delete(dbcon *sql.DB, id string) error {
	sqlstmnt := `delete from "projects" where id=$1`
	_, err := dbcon.Exec(sqlstmnt, id)
	if err != nil {
		dblogger.Println(err)
		return nil
	}
	return nil
}

// Update update a project
func Update(dbcon *sql.DB, project *Project, id string) (*Project, error) {
	sqltmnt := `UPDATE projects SET  name=$1, live_url=$2 , github=$3 WHERE id=$4`
	_, err := dbcon.Exec(sqltmnt, &project.Name, &project.LiveURL, &project.Github, id)
	if err != nil {
		dblogger.Println(err)
		return nil, err
	}
	return project, nil
}

// func CreateTb(dbcon *sql.DB) {
// 	_, err := dbcon.Exec(`CREATE TABLE IF NOT EXISTS projects (
// 		id serial PRIMARY KEY,
// 		name VARCHAR(50) NOT NULL,
// 		live_url VARCHAR(50) ,
// 		github VARCHAR(50)
// 	)`)
// 	if err != nil {
// 		log.Fatalf("could not create table REASON: %s", err)
// 	}
// }
