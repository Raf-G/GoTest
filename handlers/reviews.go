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

type ReviewHandlers struct {
	service service.ReviewsService
}

func NewReviewHandler(service service.ReviewsService) ReviewHandlers {
	return ReviewHandlers{service}
}

func (res *ReviewHandlers) AddReview(w http.ResponseWriter, r *http.Request) {
	var review domain.Review

	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newReview, err := res.service.AddReview(review)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newReview)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *ReviewHandlers) EditReview(w http.ResponseWriter, r *http.Request) {
	var review domain.Review

	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["reviewId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	review.ID = reviewID

	newReview, err := res.service.EditReview(review)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&newReview)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *ReviewHandlers) GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["reviewId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	review, err := res.service.GetReview(reviewID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&review)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *ReviewHandlers) DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["reviewId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.service.DeleteReview(reviewID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func (res *ReviewHandlers) GetReviewsProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reviews, err := res.service.GetAllReviewsProduct(productID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&reviews)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (res *ReviewHandlers) GetReviews(w http.ResponseWriter, _ *http.Request) {
	reviews, err := res.service.GetReviews()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(&reviews)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
