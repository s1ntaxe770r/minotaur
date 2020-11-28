package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/s1ntaxe770r/PPI/db"
	"gopkg.in/go-playground/validator.v9"
)

// CreateProject handles new projects additons
func CreateProject(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	var project db.Project
	v := validator.New()
	decdr := json.NewDecoder(req.Body).Decode(&project)
	if decdr != nil {
		log.Println(decdr.Error())
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "could not read request body")
	}
	err := v.Struct(project)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}
	dberr := dbcon.Create(&project).Error
	handle(dberr)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(project)

}
