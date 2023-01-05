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

func (cs *ProductService) GetOneProduct(id int) (domain.Product, error) {
	errStr := fmt.Sprintf("[services] product not fetched")
	product, err := cs.store.GetProduct(id)
	if err != nil {
		return domain.Product{}, errors.Wrap(err, errStr)
	}

	return product, nil
}

func (cs *ProductService) GetAllProducts() ([]domain.Product, error) {
	errStr := fmt.Sprintf("[services] products not fetched")
	c, err := cs.store.GetProducts()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
