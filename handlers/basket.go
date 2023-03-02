package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"example.com/m/v2/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type BasketHandlers struct {
	service service.BasketsService
}

func NewBasketHandler(service service.BasketsService) BasketHandlers {
	return BasketHandlers{service}
}

func (ch *BasketHandlers) GetBasket(w http.ResponseWriter, _ *http.Request) {
	basket, err := ch.service.GetBasket(1)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
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
	var b domain.BasketProduct

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newBasket, err := res.service.AddProductToBasket(b)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newBasket)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *BasketHandlers) DeleteProductToBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.service.DeleteProductToBasket(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func (res *BasketHandlers) DecreaseQuantityProductToBasket(w http.ResponseWriter, r *http.Request) {
	var b domain.BasketProduct

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

	b.BasketID = basketID
	b.ProductID = productID

	newProduct, err := res.service.DecreaseQuantityProductToBasket(b)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newProduct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
