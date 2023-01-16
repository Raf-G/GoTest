package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"log"
	"net/http"
)

type UserHandlers struct {
	service domain.UsersService
}

func NewUserHandler(service domain.UsersService) UserHandlers {
	return UserHandlers{service}
}

func (ih *UserHandlers) Add(w http.ResponseWriter, r *http.Request) {
	var item jsonUser

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)

		return
	}

	newItem, err := ih.service.Add(userFromJSONUser(item))
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	jsonItem := jsonUserFromUser(newItem)

	err = json.NewEncoder(w).Encode(&jsonItem)
	if err != nil {
		log.Println(err)
		return
	}
}

func (ch *UserHandlers) GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := ch.service.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	return

}
