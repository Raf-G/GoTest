package repository

import (
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"log"
)

type BasketRepository struct {
	db *sql.DB
}

func NewBasketRepository(db *sql.DB) *BasketRepository {
	return &BasketRepository{db}
}

func (rep *BasketRepository) GetBasket(userID int) (domain.Basket, error) {
	rows, err := rep.db.Query("select id from toy_shop.baskets WHERE id = ?", userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var basket domain.Basket
	var basketID int

	for rows.Next() {
		err := rows.Scan(&basketID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return basketID, nil
}
