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

func (res *ReviewService) AddReview(item domain.Review) (domain.Review, error) {
	errStr := fmt.Sprintf("[services] review not added")

	itemDB, err := res.store.AddReview(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	if itemDB == nil {
		return item, errors.Wrap(domain.ErrReviewNotFound, errStr)
	}

	return *itemDB, nil
}

func (res *ReviewService) EditReview(review domain.Review) (domain.Review, error) {
	errStr := fmt.Sprintf("[services] review not edited")

	err := res.store.EditReview(review)
	if err != nil {
		return domain.Review{}, errors.Wrap(domain.ErrReviewNotEdited, errStr)
	}

	newReview, err := res.store.GetReview(review.ID)
	if err != nil {
		return domain.Review{}, errors.Wrap(err, errStr)
	}

	return newReview, nil
}

func (res *ReviewService) GetOneReview(id int) (domain.Review, error) {
	errStr := fmt.Sprintf("[services] review not fetched")
	review, err := res.store.GetReview(id)
	if err != nil {
		return domain.Review{}, errors.Wrap(err, errStr)
	}

	return review, nil
}

func (res *ReviewService) DeleteReview(reviewID int) error {
	errStr := fmt.Sprintf("[services] review (reviewID %d) not deleted", reviewID)

	isDeleted, err := res.store.DeleteReview(reviewID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return nil
}

func (res *ReviewService) GetAllReviewsProduct(productID int) ([]domain.Review, error) {
	errStr := fmt.Sprintf("[services] reviews not fetched")

	c, err := res.store.GetReviewsProduct(productID)
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}

func (res *ReviewService) GetReviews() ([]domain.Review, error) {
	errStr := fmt.Sprintf("[services] reviews not fetched")

	c, err := res.store.GetReviews()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
