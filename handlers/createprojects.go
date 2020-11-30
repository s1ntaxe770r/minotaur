package handlers

import (
	"encoding/json"
	"io/ioutil"
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
	reqbody, err := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqbody, &project)
	if err != nil {
		log.Println(err)
		log.Println(project)
		http.Error(resp, "could not read request body", http.StatusBadRequest)
		return
	}
	vlderr := v.Struct(project)
	if vlderr != nil {
		log.Println(vlderr)
		http.Error(resp, vlderr.Error(), http.StatusBadRequest)
		return
	}
	dberr := dbcon.Create(&project).Error
	handle(dberr)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(project)
	return

}
