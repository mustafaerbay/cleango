package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mustafaerbay/cleango/internal/biz/user"
)
/*
This example shows a simple implementation of the HTTP handlers for handling CRUD operations on users. 
The UserHandler struct contains the UserService object that is used to interact with the business logic layer. 
The List, Get, Create, Update, and Delete methods correspond to the different HTTP methods and implement the corresponding functionality.

Note that this example assumes that the user.Service interface has already been defined in 
the internal/biz/user/service.go file, and that the user.User model object has already been defined 
in the internal/biz/user/model.go file. The UserHandler simply acts as an adapter between the HTTP protocol
and the business logic layer, which is implemented in the internal/biz/user package.
*/
type UserHandler struct {
	UserService user.Service
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	user, err := h.UserService.GetUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.CreateUser(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	var updateUser user.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.UpdateUser(userId, &updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if err := h.UserService.DeleteUser(userId); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
