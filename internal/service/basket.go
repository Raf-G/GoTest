package service

import (
	domain "example.com/m/v2/internal/domain"
	repository "example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/validation"
	"fmt"
	"github.com/pkg/errors"
)

type BasketsService interface {
	AddProductToBasket(domain.BasketProduct) (domain.BasketProduct, error)
	DecreaseQuantityProductToBasket(domain.BasketProduct) (domain.BasketProduct, error)
	DeleteProductToBasket(int) error
	GetBasket(int) (domain.Basket, error)
}

type BasketService struct {
	store        repository.BasketsStorage
	storeProduct repository.ProductsStorage
}

func NewBasketService(storage repository.BasketsStorage, storageProduct repository.ProductsStorage) *BasketService {
	return &BasketService{storage, storageProduct}
}

func (cs *BasketService) GetBasket(userID int) (domain.Basket, error) {
	errStr := "basket not fetched"

	var b domain.Basket

	c, err := cs.store.GetBasket(userID)
	if err != nil {
		return b, errors.Wrap(err, errStr)
	}
	b = *c
	b.UserID = userID

	// Calculation all price basket
	var totalPrice int

	for _, product := range b.Products {
		totalPrice += product.TotalPrice
	}

	b.TotalPrice = totalPrice

	return b, nil
}

func (res *BasketService) AddProductToBasket(u domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := "product to basket not added"

	err := validation.BasketProductValidation(u)
	if err != nil {
		return u, errors.Wrap(err, errStr)
	}

	existingProduct, err := res.store.GetBasketProduct(u.BasketID, u.ProductID)
	if err == nil && existingProduct != nil {
		existingProduct.Count += 1

		newProduct, errEdit := res.store.EditBasketProduct(*existingProduct)
		if errEdit != nil {
			return u, errors.Wrap(domain.ErrBasketProductNotFound, errStr)
		}

		return newProduct, nil
	}

	productDB, err := res.store.AddBasketProduct(u)
	if err != nil {
		return u, errors.Wrap(err, errStr)
	}

	errProduct := "product not fetched"
	product, err := res.storeProduct.GetProduct(u.ProductID)

	if err != nil {
		return u, errors.Wrap(err, errProduct)
	}

	productDB.TotalPrice = product.Price * productDB.Count

	return productDB, nil
}

func (res *BasketService) DecreaseQuantityProductToBasket(p domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := "product to basket not edit"
	errProduct := "product to basket not fetched"

	existingProduct, err := res.store.GetBasketProduct(p.BasketID, p.ProductID)
	if err == nil && existingProduct != nil {
		if existingProduct.Count <= 1 {
			isDeleted, errDel := res.store.DeleteBasketProduct(existingProduct.ID)
			if errDel != nil {
				return *existingProduct, errDel
			}

			if !isDeleted {
				return *existingProduct, errors.Wrap(domain.ErrBasketNotDeleted, errStr)
			}

			existingProduct.Count -= 1
			existingProduct.TotalPrice = 0
			return *existingProduct, nil
		}

		existingProduct.Count -= 1

		_, errEdit := res.store.EditBasketProduct(*existingProduct)
		if errEdit != nil {
			return p, errors.Wrap(domain.ErrBasketProductNotFound, errStr)
		}

		productInfo, errGet := res.storeProduct.GetProduct(p.ProductID)
		if errGet != nil {
			return p, errors.Wrap(errGet, errProduct)
		}

		existingProduct.TotalPrice = productInfo.Price * existingProduct.Count
		return *existingProduct, nil
	}

	p.Count -= 1
	p.TotalPrice = 0
	return p, nil
}

func (res *BasketService) DeleteProductToBasket(basketProductID int) error {
	errStr := fmt.Sprintf("product to basket (basketProductID %d) not deleted", basketProductID)

	isDeleted, err := res.store.DeleteBasketProduct(basketProductID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain.ErrBasketProductNotFound, errStr)
	}

	return nil
}
