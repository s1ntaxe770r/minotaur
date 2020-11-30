package handlers

import (
	"encoding/json"
	"fmt"
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
	encoder := json.NewEncoder(resp)
	dbcon := db.Connect()
	projects := []db.Project{}
	err := dbcon.Find(&projects).Error
	if err == nil {
		fmt.Fprintf(resp, "{}")
		return
	}
	handle(err)
	jrsp := encoder.Encode(projects)
	handle(jrsp)
	resp.Header().Set("Content-Type", "application/json")
	encoder.Encode(projects)

}
