package user

import (
	"github.com/gorilla/mux"
)

// SetRoutes add routes from user
func SetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", GetAll).Methods("GET")
	r.HandleFunc("/{name}", FindByName).Methods("GET")
}
