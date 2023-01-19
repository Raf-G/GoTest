package repository

import (
	"context"
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

func (res *BasketRepository) GetBasket(userID int) (domain.Basket, error) {
	rows, err := res.db.Query("SELECT products_baskets.id, products_baskets.basket_id, product_id, count, price * count AS total_price FROM baskets JOIN products_baskets ON baskets.id = products_baskets.basket_id JOIN products ON products_baskets.product_id = products.id WHERE user_id = ?", userID)

	if err != nil {
		log.Println(err)
		return domain.Basket{}, err
	}
	defer rows.Close()

	basket := domain.Basket{}
	var products []domain.BasketProduct

	for rows.Next() {
		p := domain.BasketProduct{}
		err = rows.Scan(&p.ID, &p.BasketID, &p.ProductID, &p.Count, &p.TotalPrice)
		if err != nil {
			fmt.Println(err)
			return domain.Basket{}, err
		}
		products = append(products, p)
	}

	basket.Products = products
	return basket, nil
}

func (res *BasketRepository) AddBasketProduct(item domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := "[repository] basket product not added to the database"

	query := "INSERT INTO `products_baskets` (`basket_id`, `product_id`, `count`) VALUES (?, ?, ?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, item.BasketID, item.ProductID, item.Count)
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	item.ID = int(id)
	log.Printf("inserted id: %d", id)

	return item, nil
}

func (res *BasketRepository) GetBasketProduct(basketID int, productID int) (domain.BasketProduct, error) {
	errStr := "[repository] basket product not fetched from the database: "

	row := res.db.QueryRow("SELECT id, basket_id, product_id, count FROM products_baskets WHERE basket_id = ? AND product_id = ?", basketID, productID)

	product := domain.BasketProduct{}

	err := row.Scan(&product.ID, &product.BasketID, &product.ProductID, &product.Count)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.BasketProduct{}, err
	}

	return product, nil
}

func (res *BasketRepository) EditBasketProduct(product domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := "[repository] basket product not edit from the database: "

	stmt, err := res.db.Prepare("UPDATE products_baskets SET count = ? WHERE id = ?")
	if err != nil {
		fmt.Println(errStr, err)
		return domain.BasketProduct{}, err
	}

	_, err = stmt.Exec(product.Count, product.ID)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.BasketProduct{}, err
	}

	return product, nil
}

func (res *BasketRepository) DeleteBasketProduct(productID int) (bool, error) {
	errStr := "[repository] basket product not deleted from the database: "

	_, err := res.db.Exec("DELETE FROM products_baskets WHERE id = ?", productID)
	if err != nil {
		fmt.Println(errStr, err)
		return false, err
	}

	return true, nil
}

func (res *BasketRepository) GetBasketProducts(basketID int) ([]domain.BasketProduct, error) {
	rows, err := res.db.Query("select * from products_baskets WHERE basket_id = ?", basketID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var basketProducts []domain.BasketProduct

	for rows.Next() {
		product := domain.BasketProduct{}
		err = rows.Scan(&product.ID, &product.BasketID, &product.ProductID, &product.Count)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		basketProducts = append(basketProducts, product)
	}
	return basketProducts, nil
}
