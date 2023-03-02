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

type ProductHandlers struct {
	service service.ProductsService
}

func NewProductHandler(service service.ProductsService) ProductHandlers {
	return ProductHandlers{service}
}

func (res *ProductHandlers) AddProduct(w http.ResponseWriter, r *http.Request) {
	var p domain.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProduct, err := res.service.AddProduct(p)
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
	var p domain.Product

	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.ID = productID

	newProduct, err := res.service.EditProduct(p)
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
