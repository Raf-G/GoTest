package service

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/repository"
	"fmt"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=products.go -destination=mocks/products.go

type ProductsService interface {
	AddProduct(domain.Product) (domain.Product, error)
	GetProduct(int) (domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(int) error
	GetAllProducts() ([]domain.Product, error)
}

type ProductService struct {
	store repository.ProductsStorage
}

func NewProductService(storage repository.ProductsStorage) *ProductService {
	return &ProductService{storage}
}

func (res *ProductService) AddProduct(p domain.Product) (domain.Product, error) {
	errStr := "product not added"

	productDB, err := res.store.AddProduct(p)
	if err != nil {
		return p, errors.Wrap(err, errStr)
	}

	if productDB == nil {
		return p, errors.Wrap(domain.ErrProductNotCreated, errStr)
	}

	return *productDB, nil
}

func (res *ProductService) GetProduct(id int) (domain.Product, error) {
	errStr := "product not fetched"
	product, err := res.store.GetProduct(id)
	if err != nil {
		return domain.Product{}, errors.Wrap(err, errStr)
	}

	return *product, nil
}

func (res *ProductService) EditProduct(p domain.Product) (domain.Product, error) {
	errStr := "product not edit"

	newProduct, err := res.store.EditProduct(p)
	if err != nil {
		return domain.Product{}, errors.Wrap(domain.ErrUserNotFound, errStr)
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
		return errors.Wrap(domain.ErrProductNotFound, errStr)
	}

	return nil
}

func (res *ProductService) GetAllProducts() ([]domain.Product, error) {
	errStr := "products not fetched"
	c, err := res.store.GetProducts()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
