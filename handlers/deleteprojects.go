package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/db"
)

// DeleteProject handles removal of projects
func DeleteProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	defer dbcon.Close()
	vars := mux.Vars(req)
	id := vars["id"]
	jrsp := db.Delete(dbcon, id)
	if jrsp != nil {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusInternalServerError)
		http.Error(resp, "could not delete project with specified id", http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}
