package routes

import (
	controllers "danspro/src/controllers/auths"
	"danspro/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func AuthenticationRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerRegister := http.HandlerFunc(controllers.ControllerStructure{}.ControllerRegister)
	r.Handle(uri(base, "/register"), ControllerRegister).Methods("POST")
	ControllerLogin := http.HandlerFunc(controllers.ControllerStructure{}.ControllerLogin)
	r.Handle(uri(base, "/login"), ControllerLogin).Methods("POST")
}
