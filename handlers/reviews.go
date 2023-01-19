package handlers

import (
	"encoding/json"
	"example.com/m/v2/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ReviewHandlers struct {
	service domain.ReviewsService
}

func NewReviewHandler(service domain.ReviewsService) ReviewHandlers {
	return ReviewHandlers{service}
}

func (res *ReviewHandlers) AddReview(w http.ResponseWriter, r *http.Request) {
	var item domain.Review

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newItem, err := res.service.AddReview(item)
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

func (res *ReviewHandlers) EditReview(w http.ResponseWriter, r *http.Request) {
	var item domain.Review

	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["reviewId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data in request body", 400)
		return
	}

	item.ID = reviewID

	newItem, err := res.service.EditReview(item)
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

func (res *ReviewHandlers) GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["reviewId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	review, err := res.service.GetOneReview(reviewID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}

	reviews, err := res.service.GetAllReviewsProduct(productID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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
		w.WriteHeader(http.StatusInternalServerError)
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
