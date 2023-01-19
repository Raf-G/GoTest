package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type BasketHandlers struct {
	service domain.BasketsService
}

func NewBasketHandler(service domain.BasketsService) BasketHandlers {
	return BasketHandlers{service}
}

func (ch *BasketHandlers) GetBasket(w http.ResponseWriter, _ *http.Request) {
	basket, err := ch.service.GetBasket(1)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(basket)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *BasketHandlers) AddProductToBasket(w http.ResponseWriter, r *http.Request) {
	var item domain.BasketProduct

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	newItem, err := res.service.AddProductToBasket(item)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	err = json.NewEncoder(w).Encode(&newItem)
	if err != nil {
		log.Println(err)
		return
	}
}

func (res *BasketHandlers) DeleteProductToBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		return
	}

	err = res.service.DeleteProductToBasket(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
}

func (res *BasketHandlers) DecreaseQuantityProductToBasket(w http.ResponseWriter, r *http.Request) {
	var item domain.BasketProduct

	vars := mux.Vars(r)

	basketID, err := strconv.Atoi(vars["basketId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	item.BasketID = basketID
	item.ProductID = productID

	newItem, err := res.service.DecreaseQuantityProductToBasket(item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newItem)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
