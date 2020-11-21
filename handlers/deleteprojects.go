package handlers

import (
	"fmt"
	"net/http"
)

// DeleteProjects handles removal of projects
func DeleteProject(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprint(resp, "DELETE")

}
