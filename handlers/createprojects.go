package handlers

import (
	"fmt"
	"net/http"
)

// CreateProjects handles new projects additons
func CreateProject(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprint(resp, "POST")

}
