package user

import (
	"github.com/gorilla/mux"
)

func GetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", GetAll).Methods("GET")
}
