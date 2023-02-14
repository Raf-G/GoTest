package repository

import (
	"context"
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

func (rep *ProductRepository) AddProduct(item domain.Product) (*domain.Product, error) {
	errStr := "[repository] product not added to the database"

	query := "INSERT INTO `products` (`name`, `description`, `price`) VALUES (?, ?, ?)"

	insertResult, err := rep.db.ExecContext(context.Background(), query, item.Name, item.Description, item.Price)
	if err != nil {
		log.Printf("%s: %s\n", errStr, err)
		return &item, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Printf("%s: %s\n", errStr, err)
		return &item, err
	}

	item.ID = int(id)
	log.Printf("inserted id: %d", id)

	return &item, nil
}

func (rep *ProductRepository) GetProduct(id int) (*domain.Product, error) {
	rows, err := rep.db.Query("select id, name, description, price from toy_shop.products WHERE products.id =?", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	product := domain.Product{}

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &product, nil
}

func (rep *ProductRepository) EditProduct(product domain.Product) (domain.Product, error) {
	errStr := "[repository] product not edit from the database: "

	stmt, err := rep.db.Prepare("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		fmt.Println(errStr, err)
		return domain.Product{}, err
	}

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.Product{}, err
	}

	return product, nil
}

func (rep *ProductRepository) DeleteProduct(productID int) (bool, error) {
	errStr := "[repository] product not deleted from the database: "

	_, err := rep.db.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
		fmt.Println(errStr, err)
		return false, err
	}

	return true, nil
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
