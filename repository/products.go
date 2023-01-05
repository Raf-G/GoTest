package repository

import (
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (rep *ProductRepository) GetProduct(id int) (domain.Product, error) {
	rows, err := rep.db.Query("select id, name, description, price from toy_shop.products WHERE products.id =?", id)
	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}
	defer rows.Close()

	product := domain.Product{}

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			fmt.Println(err)
			return domain.Product{}, err
		}
	}

	return product, nil
}

func (rep *ProductRepository) GetProducts() ([]domain.Product, error) {
	rows, err := rep.db.Query("select id, name, description, price from toy_shop.products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
