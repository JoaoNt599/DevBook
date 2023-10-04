package controllers

import (
	"api/src/base"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := base.Cannect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUsersRepository(db)
	userId, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Insert Id: %d", userId)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching all users"))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching a user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating a user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting a user"))
}
