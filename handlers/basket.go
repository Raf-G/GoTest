package handlers

import (
	"example.com/m/v2/domain"
	"html/template"
	"log"
	"net/http"
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

	tmpl, err := template.ParseFiles("templates/basket.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, basket) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
