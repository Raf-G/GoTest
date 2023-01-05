package handlers

import (
	"example.com/m/v2/domain"
	"html/template"
	"log"
	"net/http"
)

type ProductHandlers struct {
	service domain.ProductsService
}

func NewProductHandler(service domain.ProductsService) ProductHandlers {
	return ProductHandlers{service}
}

func (ch *ProductHandlers) GetProduct(w http.ResponseWriter, _ *http.Request) {
	product, err := ch.service.GetOneProduct(1)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/product.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, product) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ch *ProductHandlers) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := ch.service.GetAllProducts()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/products.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
