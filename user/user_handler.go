package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julioc98/crudgo/db"
	"github.com/julioc98/crudgo/util"
)

// GetAll Users
func GetAll(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()
	var users []User
	db.Find(&users)
	util.RespondWithJSON(w, http.StatusOK, users)
}

// FindByName find a User by name
func FindByName(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var users []User
	vars := mux.Vars(r)
	name := vars["name"]
	db.Find(&users, "name = ?", name)
	if len(users) >= 0 {
		util.RespondWithJSON(w, http.StatusOK, users)
		return
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
	db := db.Conn()
	defer db.Close()

	var user User

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
	db.Find(&user, id)
	if user.ID != 0 {
		util.RespondWithJSON(w, http.StatusOK, user)
		return
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
	db := db.Conn()
	defer db.Close()

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

	db.Create(&user)

	msg = util.Message{
		Content: "New user successfully added",
		Status:  "OK",
		Body:    user,
	}
	util.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a User by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var user User
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

	db.Find(&user, id)
	if user.ID != 0 {
		db.Delete(&user)
		msg = util.Message{
			Content: "Deleted user successfully",
			Status:  "OK",
			Body:    user,
		}

		util.RespondWithJSON(w, http.StatusAccepted, msg)
		return
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
	db := db.Conn()
	defer db.Close()

	var user User
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

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if id >= 0 {
		user.ID = uint(id)
		db.Save(&user)

		msg = util.Message{
			Content: "Update user successfully ",
			Status:  "OK",
			Body:    user,
		}
		util.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = util.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	util.RespondWithJSON(w, http.StatusOK, msg)

}
