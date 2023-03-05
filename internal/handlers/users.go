package handlers

import (
	"encoding/json"
	"example.com/m/v2/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	service service.UsersService
}

func NewUserHandler(service service.UsersService) UserHandlers {
	return UserHandlers{service}
}

func (res *UserHandlers) Add(w http.ResponseWriter, r *http.Request) {
	var u jsonUser

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err := res.service.Add(userFromJSONUser(u))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonUser := jsonUserFromUser(newUser)

	err = json.NewEncoder(w).Encode(&jsonUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// @Summary GetUser
// @Id 1
// @version 1.0
// @produce application/json
// @name Get user
// @Param user_id path int true "UserID"
// @Router /users/{user_id} [get]
// @Success 200 {object} domain.User
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

	jsonUser := jsonUserFromUser(user)

	err = json.NewEncoder(w).Encode(&jsonUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// @Summary GetUsers
// @Id 2
// @version 1.0
// @produce application/json
// @name Get users
// @Success 200 {object} []domain.User
// @Router /users [get]
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
	var u jsonUser

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.ID = userID

	newUser, err := res.service.Edit(userFromJSONUser(u))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonUser := jsonUserFromUser(newUser)

	err = json.NewEncoder(w).Encode(&jsonUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *UserHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.service.Delete(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
