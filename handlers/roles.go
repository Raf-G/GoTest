package handlers

import (
	"encoding/json"
	"example.com/m/v2/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type RoleHandlers struct {
	service service.RolesService
}

func NewRoleHandler(service service.RolesService) RoleHandlers {
	return RoleHandlers{service}
}

func (res *RoleHandlers) GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["roleId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	role, err := res.service.GetRole(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonRole := jsonRoleFromRole(role)

	err = json.NewEncoder(w).Encode(&jsonRole)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (ch *RoleHandlers) GetRoles(w http.ResponseWriter, _ *http.Request) {
	roles, err := ch.service.GetRoleAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(roles)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
