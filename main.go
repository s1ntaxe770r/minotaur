package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/PPI/handlers"
	"github.com/s1ntaxe770r/PPI/utils"
	"github.com/s1ntaxe770r/requel"
)

func main() {
	envconf := utils.NewConfig()
	envconf.LoadEnv()
	r := mux.NewRouter()
	r.Use(requel.LogReq)

	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")

	color.Green("server started on %s", envconf.ServerPort)
	log.Fatal(http.ListenAndServe(envconf.ServerPort, r))

}
