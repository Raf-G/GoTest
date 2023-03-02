package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/domain"
	"log"
)

//go:generate mockgen -source=products.go -destination=mocks/products.go

type ProductsStorage interface {
	AddProduct(domain.Product) (*domain.Product, error)
	GetProduct(int) (*domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(int) (bool, error)
	GetProducts() ([]domain.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (rep *ProductRepository) AddProduct(u domain.Product) (*domain.Product, error) {
	query := "INSERT INTO `products` (`name`, `description`, `price`) VALUES (?, ?, ?)"

	insertResult, err := rep.db.ExecContext(context.Background(), query, u.Name, u.Description, u.Price)
	if err != nil {
		return &u, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return &u, err
	}

	u.ID = int(id)

	return &u, nil
}

func (rep *ProductRepository) GetProduct(id int) (*domain.Product, error) {
	rows, err := rep.db.Query("select id, name, description, price from toy_shop.products WHERE products.id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product := domain.Product{}

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
	}

	return &product, nil
}

func (rep *ProductRepository) EditProduct(product domain.Product) (domain.Product, error) {
	stmt, err := rep.db.Prepare("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		return domain.Product{}, err
	}

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (rep *ProductRepository) DeleteProduct(productID int) (bool, error) {
	_, err := rep.db.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
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
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
