package routes

import (
	controllers "danspro/src/controllers/jobs"
	"danspro/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func GetJobsRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerGet := http.HandlerFunc(controllers.ControllerStructure{}.ControllerGetJobsList)
	r.Handle(uri(base, "/get"), ControllerGet).Methods("GET")
	ControllerGetDetail := http.HandlerFunc(controllers.ControllerStructure{}.ControllerGetJobsDetail)
	r.Handle(uri(base, "/detail/{id}"), ControllerGetDetail).Methods("GET")

}
