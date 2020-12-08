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
	json.NewDecoder(req.Body).Decode(&project)
	fmt.Println(project)
	v := validator.New()
	vlderr := v.Struct(project)
	if vlderr != nil {
		log.Println(vlderr)
		http.Error(resp, vlderr.Error(), http.StatusBadRequest)
		return
	}

	dberr := dbcon.Create(&project).Error
	handle(dberr)
	dbcon.Commit()
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(project)
	return

}
