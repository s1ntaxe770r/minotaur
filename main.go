package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/handlers"
	"github.com/s1ntaxe770r/PPI/utils"
)

func main() {
	envconf := utils.NewConfig()
	envconf.LoadEnv()
	r := mux.NewRouter()

	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")

	log.Printf("server started on  %s", envconf.ServerPort)
	log.Fatal(http.ListenAndServe(envconf.ServerPort, r))

}
