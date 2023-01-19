package service

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/validation"
	"fmt"
	"github.com/pkg/errors"
)

type BasketService struct {
	store        domain.BasketsStorage
	storeProduct domain.ProductsStorage
}

func NewBasketService(storage domain.BasketsStorage, storageProduct domain.ProductsStorage) *BasketService {
	return &BasketService{storage, storageProduct}
}

func (cs *BasketService) GetBasket(userID int) (domain.Basket, error) {
	errStr := fmt.Sprintf("[services] basket not fetched")

	var basket domain.Basket

	c, err := cs.store.GetBasket(userID)
	if err != nil {
		return basket, errors.Wrap(err, errStr)
	}
	basket = c
	basket.UserID = userID

	//Calculation all price basket
	var totalPrice int

	for _, item := range basket.Products {
		totalPrice += item.TotalPrice
	}

	basket.TotalPrice = totalPrice

	return basket, nil
}

func (res *BasketService) AddProductToBasket(item domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := fmt.Sprintf("[services] product to basket not added")

	err := validation.BasketProductValidation(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	existingProduct, err := res.store.GetBasketProduct(item.BasketID, item.ProductID)
	if err == nil {
		existingProduct.Count += 1

		newProduct, err := res.store.EditBasketProduct(existingProduct)
		if err != nil {
			return existingProduct, errors.Wrap(domain.ErrBasketProductNotFound, errStr)
		}

		return newProduct, nil
	}

	itemDB, err := res.store.AddBasketProduct(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	errProduct := fmt.Sprintf("[services] product not fetched")
	product, err := res.storeProduct.GetProduct(item.ProductID)

	if err != nil {
		return item, errors.Wrap(err, errProduct)
	}

	itemDB.TotalPrice = product.Price * itemDB.Count

	return itemDB, nil
}

func (res *BasketService) DecreaseQuantityProductToBasket(product domain.BasketProduct) (domain.BasketProduct, error) {
	errStr := fmt.Sprintf("[services] product to basket not edit")
	errProduct := fmt.Sprintf("[services] product to basket not fetched")

	existingProduct, err := res.store.GetBasketProduct(product.BasketID, product.ProductID)

	if err == nil {
		if existingProduct.Count <= 1 {
			isDeleted, err := res.store.DeleteBasketProduct(existingProduct.ID)
			if err != nil {
				return existingProduct, err
			}

			if !isDeleted {
				return existingProduct, err
			}

			existingProduct.Count -= 1
			existingProduct.TotalPrice = 0
			return existingProduct, nil
		}

		existingProduct.Count -= 1

		_, err := res.store.EditBasketProduct(existingProduct)
		if err != nil {
			return product, errors.Wrap(domain.ErrBasketProductNotFound, errStr)
		}

		productInfo, err := res.storeProduct.GetProduct(product.ProductID)
		if err != nil {
			return product, errors.Wrap(err, errProduct)
		}

		existingProduct.TotalPrice = productInfo.Price * existingProduct.Count
		return existingProduct, nil
	}

	product.Count -= 1
	product.TotalPrice = 0
	return product, nil
}

func (res *BasketService) DeleteProductToBasket(basketProductID int) error {
	errStr := fmt.Sprintf("[services] product to basket (basketProductID %d) not deleted", basketProductID)

	isDeleted, err := res.store.DeleteBasketProduct(basketProductID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain.ErrBasketProductNotFound, errStr)
	}

	return nil
}
