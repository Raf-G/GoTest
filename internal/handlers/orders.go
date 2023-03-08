package handlers

import (
	"encoding/json"
	"example.com/m/v2/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	strconv "strconv"
)

type OrderHandlers struct {
	service service.OrdersService
}

func NewOrderHandler(service service.OrdersService) OrderHandlers {
	return OrderHandlers{service}
}

// @Summary Add order
// @Tags Orders
// @produce application/json
// @Param user_id path int true "UserID"
// @Router /orders/{user_id} [post]
// @Success 200 {object} domain.Order
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

// @Summary Get order
// @Tags Orders
// @produce application/json
// @Param order_id path int true "OrderID"
// @Router /orders/{order_id} [get]
// @Success 200 {object} domain.Order
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

// @Summary Delete order
// @Tags Orders
// @produce application/json
// @Param order_id path int true "OrderID"
// @Router /orders/{order_id} [delete]
// @Success 200
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

// @Summary Get orders
// @Tags Orders
// @produce application/json
// @Router /orders [get]
// @Success 200 {object} []domain.Order
func (res *OrderHandlers) GetOrders(w http.ResponseWriter, _ *http.Request) {
	orders, err := res.service.GetOrders()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
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
