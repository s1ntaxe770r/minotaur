package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/s1ntaxe770r/minotaur/db"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetProjects returns all available projects
func GetProjects(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	defer dbcon.Close()
	resp.Header().Set("Content-Type", "application/json")
	projects, err := db.QueryAll(dbcon)
	if err != nil {
		log.Println(err)
		http.Error(resp, "could not retrieve projects from the db", http.StatusInternalServerError)
		return
	}
	handle(err)
	jrsp, err := json.Marshal(projects)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jrsp)
}
