package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ProductHandlers struct {
	service domain.ProductsService
}

func NewProductHandler(service domain.ProductsService) ProductHandlers {
	return ProductHandlers{service}
}

func (res *ProductHandlers) AddProduct(w http.ResponseWriter, r *http.Request) {
	var item domain.Product

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	newItem, err := res.service.AddProduct(item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newItem)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (ch *ProductHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := ch.service.GetProduct(productID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (ch *ProductHandlers) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := ch.service.GetAllProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&products)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *ProductHandlers) EditProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product

	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	product.ID = productID

	newItem, err := res.service.EditProduct(product)
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

func (res *ProductHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.service.DeleteProduct(productID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
