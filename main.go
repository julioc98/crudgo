package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/julioc98/crudgo/user"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := mux.NewRouter()

	user.GetRoutes(r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handler)
	http.Handle("/", r)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
