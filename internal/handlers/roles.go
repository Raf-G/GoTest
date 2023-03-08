package handlers

import (
	"encoding/json"
	"example.com/m/v2/internal/service"
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

// @Summary Get role
// @Tags Roles
// @produce application/json
// @Param role_id path int true "RoleID"
// @Router /roles/{role_id} [get]
// @Success 200 {object} domain.Role
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

// @Summary Get roles
// @Tags Roles
// @produce application/json
// @Router /roles [get]
// @Success 200 {object} []domain.Role
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
