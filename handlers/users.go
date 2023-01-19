package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	service domain.UsersService
}

func NewUserHandler(service domain.UsersService) UserHandlers {
	return UserHandlers{service}
}

func (res *UserHandlers) Add(w http.ResponseWriter, r *http.Request) {
	var item jsonUser

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	newItem, err := res.service.Add(userFromJSONUser(item))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonItem := jsonUserFromUser(newItem)

	err = json.NewEncoder(w).Encode(&jsonItem)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *UserHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := res.service.GetUser(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonItem := jsonUserFromUser(user)

	err = json.NewEncoder(w).Encode(&jsonItem)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *UserHandlers) GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := res.service.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *UserHandlers) Edit(w http.ResponseWriter, r *http.Request) {
	var item jsonUser

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	item.ID = userID

	newItem, err := res.service.Edit(userFromJSONUser(item))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonItem := jsonUserFromUser(newItem)

	err = json.NewEncoder(w).Encode(&jsonItem)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *UserHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	err := res.service.Delete(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
