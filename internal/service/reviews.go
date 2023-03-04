package service

import (
	domain2 "example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repository"
	"fmt"
	"github.com/pkg/errors"
)

type ReviewsService interface {
	AddReview(domain2.Review) (domain2.Review, error)
	EditReview(domain2.Review) (domain2.Review, error)
	GetReview(int) (domain2.Review, error)
	DeleteReview(int) error
	GetAllReviewsProduct(int) ([]domain2.Review, error)
	GetReviews() ([]domain2.Review, error)
}

type ReviewService struct {
	store repository.ReviewsStorage
}

func NewReviewService(storage repository.ReviewsStorage) *ReviewService {
	return &ReviewService{storage}
}

func (res *ReviewService) AddReview(r domain2.Review) (domain2.Review, error) {
	errStr := "review not added"

	reviewDB, err := res.store.AddReview(r)
	if err != nil {
		return r, errors.Wrap(err, errStr)
	}

	if reviewDB == nil {
		return r, errors.Wrap(domain2.ErrReviewNotFound, errStr)
	}

	return *reviewDB, nil
}

func (res *ReviewService) EditReview(r domain2.Review) (domain2.Review, error) {
	errStr := "review not edited"

	err := res.store.EditReview(r)
	if err != nil {
		return domain2.Review{}, errors.Wrap(domain2.ErrReviewNotEdited, errStr)
	}

	newReview, err := res.store.GetReview(r.ID)
	if err != nil {
		return domain2.Review{}, errors.Wrap(err, errStr)
	}

	return newReview, nil
}

func (res *ReviewService) GetReview(id int) (domain2.Review, error) {
	errStr := "review not fetched"
	review, err := res.store.GetReview(id)
	if err != nil {
		return domain2.Review{}, errors.Wrap(err, errStr)
	}

	return review, nil
}

func (res *ReviewService) DeleteReview(reviewID int) error {
	errStr := fmt.Sprintf("review (reviewID %d) not deleted", reviewID)

	isDeleted, err := res.store.DeleteReview(reviewID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	return nil
}

func (res *ReviewService) GetAllReviewsProduct(productID int) ([]domain2.Review, error) {
	errStr := "reviews not fetched"

	c, err := res.store.GetReviewsProduct(productID)
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}

func (res *ReviewService) GetReviews() ([]domain2.Review, error) {
	errStr := "reviews not fetched"

	c, err := res.store.GetReviews()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
