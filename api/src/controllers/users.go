package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating user"))
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
