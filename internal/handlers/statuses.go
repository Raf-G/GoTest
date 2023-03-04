package handlers

import (
	"encoding/json"
	"example.com/m/v2/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type StatusHandlers struct {
	service service.StatusesService
}

func NewStatusHandler(service service.StatusesService) StatusHandlers {
	return StatusHandlers{service}
}

func (res *StatusHandlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	statusID, err := strconv.Atoi(vars["statusId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, err := res.service.GetStatus(statusID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *StatusHandlers) GetStatuses(w http.ResponseWriter, _ *http.Request) {
	statuses, err := res.service.GetStatuses()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(statuses)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
