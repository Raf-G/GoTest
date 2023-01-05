package handlers

import (
	"example.com/m/v2/domain"
	"html/template"
	"log"
	"net/http"
)

type RoleHandlers struct {
	service domain.RolesService
}

func NewRoleHandler(service domain.RolesService) RoleHandlers {
	return RoleHandlers{service}
}

func (ch *RoleHandlers) GetRoles(w http.ResponseWriter, _ *http.Request) {
	users, err := ch.service.GetRoleAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/roles.html")
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
