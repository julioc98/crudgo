package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/crudgo/util"
)

// GetAll Users
func GetAll(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, Users)
}

// Find a User by name
func FindByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for _, user := range Users {
		if name == user.Name {
			util.RespondWithJSON(w, http.StatusOK, user)
			return
		}

	}

	msg := util.Message{
		Content: "Not exist this user",
		Status:  "ERRO",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a User
func Add(w http.ResponseWriter, r *http.Request) {
	var user User
	var msg util.Message

	msg = util.Message{
		Content: "Invalid request payload",
		Status:  "ERRO",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	Users = append(Users, user)

	msg = util.Message{
		Content: "Novo Usuario adcionado com sucesso",
		Status:  "OK",
		Body:    user,
	}
	util.RespondWithJSON(w, http.StatusCreated, msg)
}
