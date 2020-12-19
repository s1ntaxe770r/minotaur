package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/db"
)

// UpdateProject handles project updates
func UpdateProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	vars := mux.Vars(req)
	id := vars["id"]
	var project db.Project
	json.NewDecoder(req.Body).Decode(&project)
	newproject, err := db.Update(dbcon, &project, id)
	handle(err)
	json.NewEncoder(resp).Encode(&newproject)
	return

}
