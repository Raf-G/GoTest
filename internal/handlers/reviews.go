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

type ReviewHandlers struct {
	service service.ReviewsService
}

func NewReviewHandler(service service.ReviewsService) ReviewHandlers {
	return ReviewHandlers{service}
}

// @Summary Add review
// @Tags Reviews
// @produce application/json
// @Param review body domain.Review true "new review"
// @Router /reviews [post]
// @Success 200 {object} domain.Review
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

// @Summary Edit review
// @Tags Reviews
// @produce application/json
// @Param review_id path int true "ReviewID"
// @Param review body domain.Review true "new review"
// @Router /reviews/{review_id} [put]
// @Success 200 {object} domain.Review
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

// @Summary Get review
// @Tags Reviews
// @produce application/json
// @Param review_id path int true "ReviewID"
// @Router /reviews/{review_id} [get]
// @Success 200 {object} domain.Review
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

// @Summary Delete review
// @Tags Reviews
// @produce application/json
// @Param review_id path int true "ReviewID"
// @Router /reviews/{review_id} [delete]
// @Success 200
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

// @Summary Get reviews product
// @Tags Reviews
// @produce application/json
// @Param product_id path int true "ProductID"
// @Router /reviews/product/{product_id} [get]
// @Success 200 {object} []domain.Review
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

// @Summary Get reviews
// @Tags Reviews
// @produce application/json
// @Router /reviews [get]
// @Success 200 {object} []domain.Review
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
