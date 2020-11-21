package handlers

import (
	"net/http"
)

// GetProjects returns all available projects
func GetProjects(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	// dbcon := db.Connect()
	// var project db.Project

}
