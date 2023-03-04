package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/internal/domain"
)

//go:generate mockgen -source=baskets.go -destination=mocks/baskets.go

type BasketsStorage interface {
	AddBasketProduct(domain.BasketProduct) (domain.BasketProduct, error)
	GetBasketProduct(int, int) (*domain.BasketProduct, error)
	GetBasketProducts(int) ([]domain.BasketProduct, error)
	EditBasketProduct(domain.BasketProduct) (domain.BasketProduct, error)
	DeleteBasketProduct(int) (bool, error)
	GetBasket(int) (*domain.Basket, error)
}

type BasketRepository struct {
	db *sql.DB
}

func NewBasketRepository(db *sql.DB) *BasketRepository {
	return &BasketRepository{db}
}

func (res *BasketRepository) GetBasket(userID int) (*domain.Basket, error) {
	rows, err := res.db.Query("SELECT products_baskets.id, products_baskets.basket_id, product_id, count, price * count AS total_price FROM baskets JOIN products_baskets ON baskets.id = products_baskets.basket_id JOIN products ON products_baskets.product_id = products.id WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	basket := domain.Basket{}
	var products []domain.BasketProduct

	for rows.Next() {
		p := domain.BasketProduct{}
		err = rows.Scan(&p.ID, &p.BasketID, &p.ProductID, &p.Count, &p.TotalPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	basket.Products = products
	return &basket, nil
}

func (res *BasketRepository) AddBasketProduct(u domain.BasketProduct) (domain.BasketProduct, error) {
	query := "INSERT INTO `products_baskets` (`basket_id`, `product_id`, `count`) VALUES (?, ?, ?)"

	insertResult, err := res.db.ExecContext(context.Background(), query, u.BasketID, u.ProductID, u.Count)
	if err != nil {
		return u, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = int(id)

	return u, nil
}

func (res *BasketRepository) GetBasketProduct(basketID int, productID int) (*domain.BasketProduct, error) {
	row := res.db.QueryRow("SELECT id, basket_id, product_id, count FROM products_baskets WHERE basket_id = ? AND product_id = ?", basketID, productID)

	product := domain.BasketProduct{}

	err := row.Scan(&product.ID, &product.BasketID, &product.ProductID, &product.Count)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (res *BasketRepository) EditBasketProduct(product domain.BasketProduct) (domain.BasketProduct, error) {
	stmt, err := res.db.Prepare("UPDATE products_baskets SET count = ? WHERE id = ?")
	if err != nil {
		return domain.BasketProduct{}, err
	}

	_, err = stmt.Exec(product.Count, product.ID)
	if err != nil {
		return domain.BasketProduct{}, err
	}

	return product, nil
}

func (res *BasketRepository) DeleteBasketProduct(productID int) (bool, error) {
	_, err := res.db.Exec("DELETE FROM products_baskets WHERE id = ?", productID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (res *BasketRepository) GetBasketProducts(basketID int) ([]domain.BasketProduct, error) {
	rows, err := res.db.Query("select * from products_baskets WHERE basket_id = ?", basketID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var basketProducts []domain.BasketProduct

	for rows.Next() {
		product := domain.BasketProduct{}
		err = rows.Scan(&product.ID, &product.BasketID, &product.ProductID, &product.Count)
		if err != nil {
			return nil, err
		}
		basketProducts = append(basketProducts, product)
	}
	return basketProducts, nil
}
