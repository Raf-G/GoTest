package handlers

import (
	"example.com/m/v2/domain"
	"html/template"
	"log"
	"net/http"
)

type UserHandlers struct {
	service domain.UsersService
}

func NewUserHandler(service domain.UsersService) UserHandlers {
	return UserHandlers{service}
}

func (ch *UserHandlers) GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := ch.service.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, users) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
