package routes

import (
	"github.com/gorilla/mux"
)

// Route is as
func Route() *mux.Router {

	r := mux.NewRouter()
	AuthenticationRoutes(r, "/auth")
	GetJobsRoutes(r, "/jobs")
	return r
}
