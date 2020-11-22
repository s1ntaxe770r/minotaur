package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/db"
)

// DeleteProject handles removal of projects
func DeleteProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Printf("id = %s", id)
	var project db.Project
	jrsp := dbcon.Delete(&project, "id = ?", id).Error
	if jrsp != nil {
		ex := jrsp.Error()
		log.Println(ex)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusInternalServerError)
		http.Error(resp, "could not delete project with specified id", http.StatusInternalServerError)
		log.Println(jrsp)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}
