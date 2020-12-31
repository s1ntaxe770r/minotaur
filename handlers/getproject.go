package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/minotaur/db"
)

// GetProject returns the project with the specified id
func GetProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	defer dbcon.Close()
	vars := mux.Vars(req)
	id := vars["id"]
	var project db.Project
	jrsp := db.QueryOne(dbcon, &project, id)
	if jrsp != nil {
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "failed to retrieve project with id %s ", id)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(project)

}
