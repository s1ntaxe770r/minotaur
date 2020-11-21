package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/db"
)

// GetProject returns the project with the specified id
func GetProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	vars := mux.Vars(req)
	id := vars["id"]
	var project db.Project
	jrsp := dbcon.First(&project, "id = ?", id).Error
	if jrsp != nil {
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "failed to retrieve project with id %s ", id)
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(jrsp)

}
