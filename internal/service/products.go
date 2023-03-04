package service

import (
	domain2 "example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repository"
	"fmt"
	"github.com/pkg/errors"
)

type ProductsService interface {
	AddProduct(domain2.Product) (domain2.Product, error)
	GetProduct(int) (domain2.Product, error)
	EditProduct(domain2.Product) (domain2.Product, error)
	DeleteProduct(int) error
	GetAllProducts() ([]domain2.Product, error)
}

type ProductService struct {
	store repository.ProductsStorage
}

func NewProductService(storage repository.ProductsStorage) *ProductService {
	return &ProductService{storage}
}

func (res *ProductService) AddProduct(p domain2.Product) (domain2.Product, error) {
	errStr := "product not added"

	productDB, err := res.store.AddProduct(p)
	if err != nil {
		return p, errors.Wrap(err, errStr)
	}

	if productDB == nil {
		return p, errors.Wrap(domain2.ErrProductNotCreated, errStr)
	}

	return *productDB, nil
}

func (res *ProductService) GetProduct(id int) (domain2.Product, error) {
	errStr := "product not fetched"
	product, err := res.store.GetProduct(id)
	if err != nil {
		return domain2.Product{}, errors.Wrap(err, errStr)
	}

	return *product, nil
}

func (res *ProductService) EditProduct(p domain2.Product) (domain2.Product, error) {
	errStr := "product not edit"

	newProduct, err := res.store.EditProduct(p)
	if err != nil {
		return domain2.Product{}, errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	return newProduct, nil
}

func (res *ProductService) DeleteProduct(productID int) error {
	errStr := fmt.Sprintf("product (productID %d) not deleted", productID)

	isDeleted, err := res.store.DeleteProduct(productID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain2.ErrProductNotFound, errStr)
	}

	return nil
}

func (res *ProductService) GetAllProducts() ([]domain2.Product, error) {
	errStr := "products not fetched"
	c, err := res.store.GetProducts()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
