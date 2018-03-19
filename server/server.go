package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/crudgo/logs"
	"github.com/julioc98/crudgo/user"
	"github.com/julioc98/crudgo/util"
)

var name string

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao CRUDGo @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := util.GetOSEnvironment("PORT", "5001")

	name = util.GetOSEnvironment("NAME", "JC")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)

	user.SetRoutes(r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handler)
	http.Handle("/", r)

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
