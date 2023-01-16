package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
)

type ReviewService struct {
	store domain.ReviewsStorage
}

func NewReviewService(storage domain.ReviewsStorage) *ReviewService {
	return &ReviewService{storage}
}

func (cs *ReviewService) GetOneReview(id int) (domain.Review, error) {
	errStr := fmt.Sprintf("[services] review not fetched")
	review, err := cs.store.GetReview(id)
	if err != nil {
		return domain.Review{}, errors.Wrap(err, errStr)
	}

	return review, nil
}

func (cs *ReviewService) GetAllReviewsProduct(productID int) ([]domain.Review, error) {
	errStr := fmt.Sprintf("[services] products not fetched")
	c, err := cs.store.GetReviewsProduct(productID)
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
