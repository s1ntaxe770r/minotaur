package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/s1ntaxe770r/PPI/db"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetProjects returns all available projects
func GetProjects(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	dbcon := db.Connect()
	var project db.Project
	err := dbcon.Find(&project).Error
	handle(err)
	jrsp := json.NewEncoder(resp).Encode(project)
	handle(jrsp)
	json.NewEncoder(resp).Encode(project)

}
