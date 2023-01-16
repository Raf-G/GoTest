package handlers

import (
	"example.com/m/v2/domain"
	"html/template"
	"log"
	"net/http"
)

type ReviewHandlers struct {
	service domain.ReviewsService
}

func NewReviewHandler(service domain.ReviewsService) ReviewHandlers {
	return ReviewHandlers{service}
}

func (ch *ReviewHandlers) GetReview(w http.ResponseWriter, _ *http.Request) {
	review, err := ch.service.GetOneReview(1)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/review.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, review) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ch *ReviewHandlers) GetReviewsProduct(w http.ResponseWriter, _ *http.Request) {
	reviews, err := ch.service.GetAllReviewsProduct(1)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/reviews.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, reviews) // tmpl.Execute write WriteHeader 200
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
