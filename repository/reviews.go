package repository

import (
	"context"
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

func (rep *ReviewRepository) AddReview(item domain.Review) (*domain.Review, error) {
	errStr := "[repository] review not added to the database"

	query := "INSERT INTO `reviews` (`user_id`, `product_id`, `description`, `grade`) VALUES (?, ?, ?, ?)"
	insertResult, err := rep.db.ExecContext(context.Background(), query, item.UserID, item.ProductID, item.Description, item.Grade)
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return &item, err
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return &item, err
	}
	item.ID = int(id)
	log.Printf("inserted id: %d", id)

	return &item, nil
}

func (rep *ReviewRepository) EditReview(review domain.Review) error {
	errStr := "[repository] review not edit from the database: "

	stmt, err := rep.db.Prepare("UPDATE reviews SET description = ?, grade = ? WHERE id = ?")
	if err != nil {
		fmt.Println(errStr, err)
		return err
	}

	_, err = stmt.Exec(review.Description, review.Grade, review.ID)
	if err != nil {
		fmt.Println(errStr, err)
		return err
	}

	return nil
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

func (rep *ReviewRepository) DeleteReview(reviewId int) (bool, error) {
	errStr := "[repository] review not deleted from the database: "

	_, err := rep.db.Exec("DELETE FROM reviews WHERE id = ?", reviewId)
	if err != nil {
		fmt.Println(errStr, err)
		return false, err
	}

	return true, nil
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

func (rep *ReviewRepository) GetReviews() ([]domain.Review, error) {
	rows, err := rep.db.Query("select id, user_id, product_id, description, grade from reviews")
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
