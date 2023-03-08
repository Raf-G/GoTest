package handlers

import (
	"encoding/json"
	"example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/service"
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

// @Summary Get basket
// @Tags Baskets
// @produce application/json
// @Param basket_id path int true "BasketID"
// @Router /baskets/{basket_id} [get]
// @Success 200 {object} domain.Basket
func (ch *BasketHandlers) GetBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["basketId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	basket, err := ch.service.GetBasket(id)
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

// @Summary Add product to basket
// @Tags Baskets
// @produce application/json
// @Param product body domain.BasketProduct true "new product added"
// @Router /baskets/product [post]
// @Success 200 {object} domain.BasketProduct
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

// @Summary Delete product to basket
// @Tags Baskets
// @produce application/json
// @Param basket_id path int true "BasketID"
// @Router /baskets/product/{basket_id} [delete]
// @Success 200
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

// @Summary Decrease quantity product to basket
// @Tags Baskets
// @produce application/json
// @Param product_id path int true "ProductID"
// @Param basket_id path int true "BasketID"
// @Param product body domain.BasketProduct true "new product added"
// @Router /baskets/product/{product_id}/{basket_id} [put]
// @Success 200 {object} domain.BasketProduct
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
