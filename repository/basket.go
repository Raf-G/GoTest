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

	rows, err := rep.db.Query("SELECT product_id, count, price * count AS total_price FROM baskets JOIN products_baskets ON baskets.id = products_baskets.basket_id JOIN products ON products_baskets.product_id = products.id WHERE user_id = ?", userID)

	if err != nil {
		log.Println(err)
		return domain.Basket{}, err
	}
	defer rows.Close()

	basket := domain.Basket{}
	var products []domain.BasketProduct

	for rows.Next() {
		p := domain.BasketProduct{}
		err := rows.Scan(&p.ProductID, &p.Count, &p.TotalPrice)
		if err != nil {
			fmt.Println(err)
			return domain.Basket{}, err
		}
		products = append(products, p)
	}

	basket.Products = products
	return basket, nil
}
