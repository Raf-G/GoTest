package service

import (
	domain2 "example.com/m/v2/internal/domain"
	repository2 "example.com/m/v2/internal/repository"
	"fmt"
	"github.com/pkg/errors"
)

type OrdersService interface {
	AddOrder(int) (domain2.Order, error)
	GetOrder(int) (domain2.Order, error)
	DeleteOrder(int) error
	GetOrders() ([]domain2.Order, error)
}

type OrderService struct {
	store        repository2.OrdersStorage
	storeBasket  repository2.BasketsStorage
	storeProduct repository2.ProductsStorage
}

func NewOrderService(storage repository2.OrdersStorage, storageBasket repository2.BasketsStorage, storageProduct repository2.ProductsStorage) *OrderService {
	return &OrderService{storage, storageBasket, storageProduct}
}

func (res *OrderService) AddOrder(userID int) (domain2.Order, error) {
	errStr := " order not added"

	_, err := res.storeBasket.GetBasket(userID)
	if err != nil {
		return domain2.Order{}, errors.Wrap(err, errStr)
	}

	basketProducts, err := res.storeBasket.GetBasketProducts(userID)
	if err != nil {
		return domain2.Order{}, errors.Wrap(err, errStr)
	}

	if len(basketProducts) == 0 {
		return domain2.Order{}, errors.Wrap(domain2.ErrBasketEmpty, errStr)
	}

	var productsOrder []domain2.ProductOrderForAddOrder

	for _, v := range basketProducts {
		productOrder := domain2.ProductOrderForAddOrder{}
		productOrder.ProductID = v.ProductID
		productOrder.Count = v.Count
		productOrder.BasketProductID = v.ID

		product, errGet := res.storeProduct.GetProduct(v.ProductID)
		if errGet != nil {
			return domain2.Order{}, errors.Wrap(errGet, errStr)
		}

		productOrder.Price = v.Count * product.Price

		productsOrder = append(productsOrder, productOrder)
	}

	orderID, err := res.store.AddOrder(userID)
	if err != nil {
		return domain2.Order{}, errors.Wrap(err, errStr)
	}

	if orderID == 0 {
		return domain2.Order{}, errors.Wrap(domain2.ErrOrderNotCreated, errStr)
	}

	newOrder, err := res.store.GetOrder(orderID)
	if err != nil {
		return domain2.Order{}, errors.Wrap(err, errStr)
	}

	for _, v := range productsOrder {
		productOrder := domain2.ProductOrder{}
		productOrder.OrderID = orderID
		productOrder.ProductID = v.ProductID
		productOrder.Count = v.Count
		productOrder.Price = v.Price

		_, err = res.store.AddProductOrder(productOrder)
		if err != nil {
			return domain2.Order{}, errors.Wrap(err, errStr)
		}

		isDeleted, errDel := res.storeBasket.DeleteBasketProduct(v.BasketProductID)
		if errDel != nil {
			return domain2.Order{}, errors.Wrap(errDel, errStr)
		}
		if !isDeleted {
			return domain2.Order{}, errors.Wrap(errDel, errStr)
		}
	}

	return newOrder, nil
}

func (res *OrderService) GetOrder(orderID int) (domain2.Order, error) {
	errStr := fmt.Sprintf("order (orderID %d) not fetched", orderID)

	order, err := res.store.GetOrder(orderID)
	if err != nil {
		return domain2.Order{}, errors.Wrap(err, errStr)
	}

	return order, err
}

func (res *OrderService) DeleteOrder(orderID int) error {
	errStr := fmt.Sprintf("order (orderID %d) not deleted", orderID)

	isDeleted, err := res.store.DeleteProductsOrder(orderID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	isDeleted, err = res.store.DeleteOrder(orderID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	return nil
}

func (res *OrderService) GetOrders() ([]domain2.Order, error) {
	errStr := "orders not fetched"
	c, err := res.store.GetOrders()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
