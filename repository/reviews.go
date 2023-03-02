package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/domain"
)

//go:generate mockgen -source=reviews.go -destination=mocks/reviews.go

type ReviewsStorage interface {
	AddReview(domain.Review) (*domain.Review, error)
	EditReview(domain.Review) error
	GetReview(int) (domain.Review, error)
	DeleteReview(int) (bool, error)
	GetReviewsProduct(int) ([]domain.Review, error)
	GetReviews() ([]domain.Review, error)
}

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (rep *ReviewRepository) AddReview(r domain.Review) (*domain.Review, error) {
	query := "INSERT INTO `reviews` (`user_id`, `product_id`, `description`, `grade`) VALUES (?, ?, ?, ?)"
	insertResult, err := rep.db.ExecContext(context.Background(), query, r.UserID, r.ProductID, r.Description, r.Grade)
	if err != nil {
		return &r, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return &r, err
	}

	r.ID = int(id)

	return &r, nil
}

func (rep *ReviewRepository) EditReview(r domain.Review) error {
	stmt, err := rep.db.Prepare("UPDATE reviews SET description = ?, grade = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(r.Description, r.Grade, r.ID)
	if err != nil {
		return err
	}

	return nil
}
func (rep *ReviewRepository) GetReview(id int) (domain.Review, error) {
	row := rep.db.QueryRow("select id, user_id, product_id, description, grade from reviews WHERE reviews.id =?", id)

	r := domain.Review{}

	err := row.Scan(&r.ID, &r.UserID, &r.ProductID, &r.Description, &r.Grade)
	if err != nil {
		return domain.Review{}, err
	}

	return r, nil
}

func (rep *ReviewRepository) DeleteReview(reviewId int) (bool, error) {
	_, err := rep.db.Exec("DELETE FROM reviews WHERE id = ?", reviewId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (rep *ReviewRepository) GetReviewsProduct(productID int) ([]domain.Review, error) {
	rows, err := rep.db.Query("select id, user_id, product_id, description, grade from reviews WHERE reviews.product_id =?", productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var r []domain.Review

	for rows.Next() {
		review := domain.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.ProductID, &review.Description, &review.Grade)
		if err != nil {
			return nil, err
		}
		r = append(r, review)
	}

	return r, nil
}

func (rep *ReviewRepository) GetReviews() ([]domain.Review, error) {
	rows, err := rep.db.Query("select id, user_id, product_id, description, grade from reviews")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var r []domain.Review

	for rows.Next() {
		review := domain.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.ProductID, &review.Description, &review.Grade)
		if err != nil {
			return nil, err
		}
		r = append(r, review)
	}
	return r, nil
}
