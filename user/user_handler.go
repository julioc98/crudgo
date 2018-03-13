package user

import (
	"encoding/json"
	"net/http"

	"github.com/julioc98/crudgo/util"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, Users)
}

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
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusCreated, msg)
}
