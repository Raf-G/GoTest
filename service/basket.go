package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
)

type BasketService struct {
	store domain.BasketsStorage
}

func NewBasketService(storage domain.BasketsStorage) *BasketService {
	return &BasketService{storage}
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
