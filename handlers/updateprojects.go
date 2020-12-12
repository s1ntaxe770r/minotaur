package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/db"
)

// UpdateProjects handles project updates

func UpdateProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	vars := mux.Vars(req)
	id := vars["id"]
	var project db.Project
	json.NewDecoder(req.Body).Decode(&project)
	qry := dbcon.First(&project, "id = ?", id)
	err := qry.Save(&project).Error
	handle(err)
	json.NewEncoder(resp).Encode(project)
	return

}
