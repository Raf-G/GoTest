package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
)

type ProductService struct {
	store domain.ProductsStorage
}

func NewProductService(storage domain.ProductsStorage) *ProductService {
	return &ProductService{storage}
}

func (res *ProductService) AddProduct(item domain.Product) (domain.Product, error) {
	errStr := "[services] product not added"

	itemDB, err := res.store.AddProduct(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	if itemDB == nil {
		return item, errors.Wrap(domain.ErrProductNotCreated, errStr)
	}

	return *itemDB, nil
}

func (res *ProductService) GetProduct(id int) (domain.Product, error) {
	errStr := "[services] product not fetched"
	product, err := res.store.GetProduct(id)
	if err != nil {
		return domain.Product{}, errors.Wrap(err, errStr)
	}

	return *product, nil
}

func (res *ProductService) EditProduct(product domain.Product) (domain.Product, error) {
	errStr := "[services] product not edit"

	newProduct, err := res.store.EditProduct(product)
	if err != nil {
		return domain.Product{}, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return newProduct, nil
}

func (res *ProductService) DeleteProduct(productID int) error {
	errStr := fmt.Sprintf("[services] product (productID %d) not deleted", productID)

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
	errStr := "[services] products not fetched"
	c, err := res.store.GetProducts()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
