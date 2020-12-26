package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/s1ntaxe770r/PPI/db"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetProjects returns all available projects
func GetProjects(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	defer dbcon.Close()
	resp.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(resp)
	var projects db.Projects
	err := db.QueryAll(dbcon, &projects)
	print(err)
	if err != nil {
		log.Println(err)
		return
	}
	if err == nil {
		resp.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(resp, "no entries found pls update the api")
		return
	}
	handle(err)
	jrsp := encoder.Encode(projects)
	handle(jrsp)
	resp.Header().Set("Content-Type", "application/json")
	encoder.Encode(jrsp)

}
