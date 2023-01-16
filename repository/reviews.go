package repository

import (
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"log"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (rep *ReviewRepository) GetReview(id int) (domain.Review, error) {
	row := rep.db.QueryRow("select id, user_id, product_id, description, grade from reviews WHERE reviews.id =?", id)

	review := domain.Review{}

	err := row.Scan(&review.ID, &review.UserID, &review.ProductID, &review.Description, &review.Grade)
	if err != nil {
		fmt.Println(err)
		return domain.Review{}, err
	}

	return review, nil
}

func (rep *ReviewRepository) GetReviewsProduct(productID int) ([]domain.Review, error) {
	rows, err := rep.db.Query("select id, user_id, product_id, description, grade from reviews WHERE reviews.product_id =?", productID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var reviews []domain.Review

	for rows.Next() {
		review := domain.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.ProductID, &review.Description, &review.Grade)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
