package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	strconv "strconv"
)

type OrderHandlers struct {
	service domain.OrdersService
}

func NewOrderHandler(service domain.OrdersService) OrderHandlers {
	return OrderHandlers{service}
}

func (res *OrderHandlers) AddOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order, err := res.service.AddOrder(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *OrderHandlers) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderID, err := strconv.Atoi(vars["orderId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order, err := res.service.GetOrder(orderID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&order)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *OrderHandlers) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderID, err := strconv.Atoi(vars["orderId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.service.DeleteOrder(orderID)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
}

func (res *OrderHandlers) GetOrders(w http.ResponseWriter, _ *http.Request) {
	orders, err := res.service.GetOrders()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
