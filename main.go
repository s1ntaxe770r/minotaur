package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/handlers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")

	log.Println("listening on :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
