package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julioc98/crudgo/util"
)

// GetAll Users
func GetAll(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, Users)
}

// FindByName find a User by name
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
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// FindByID find a User by ID
func FindByID(w http.ResponseWriter, r *http.Request) {

	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	for _, user := range Users {
		if id == user.ID {
			util.RespondWithJSON(w, http.StatusOK, user)
			return
		}

	}

	msg = util.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
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
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	IDBase = IDBase + 1

	user.ID = IDBase

	Users = append(Users, user)

	msg = util.Message{
		Content: "New user successfully added",
		Status:  "OK",
		Body:    user,
	}
	util.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a User by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {

	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}
	msg = util.Message{
		Content: "Deleted user successfully",
		Status:  "OK",
		Body:    nil,
	}

	for i, user := range Users {
		if id == user.ID {

			Users[i] = Users[len(Users)-1]
			Users = Users[:len(Users)-1]

			util.RespondWithJSON(w, http.StatusAccepted, msg)
			return
		}

	}

	msg = util.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a User by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	var myUser User
	var msg util.Message

	msg = util.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	msg = util.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&myUser); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	for i := 0; i < len(Users); i++ {
		if id == Users[i].ID {

			Users[i].Name = myUser.Name
			Users[i].Age = myUser.Age

			util.RespondWithJSON(w, http.StatusAccepted, Users[i])
			return
		}
	}

	msg = util.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}
