package handlers

import (
	"fmt"
	"net/http"
)

// UpdateProjects handles project updates

func UpdateProject(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprint(resp, "PUT")

}
