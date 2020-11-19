package main

import (
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/projects", handlers.getprojects).Methods("GET")
	r.HandleFunc("/projects", handlers.createproject).Methods("POST")
	r.HandleFunc("/projects{id}", handlers.deleteproject).Methods("DELETE")
	r.HandleFunc("/projects{id}", handlers.updateproject).Methods("PUT")

}
