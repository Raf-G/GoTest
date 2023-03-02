package service

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/repository"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

//go:generate mockgen -source=orders.go -destination=mocks/orders.go

type OrdersService interface {
	AddOrder(int) (domain.Order, error)
	GetOrder(int) (domain.Order, error)
	DeleteOrder(int) error
	GetOrders() ([]domain.Order, error)
}

type OrderService struct {
	store        repository.OrdersStorage
	storeBasket  repository.BasketsStorage
	storeProduct repository.ProductsStorage
}

func NewOrderService(storage repository.OrdersStorage, storageBasket repository.BasketsStorage, storageProduct repository.ProductsStorage) *OrderService {
	return &OrderService{storage, storageBasket, storageProduct}
}

func (res *OrderService) AddOrder(userID int) (domain.Order, error) {
	errStr := " order not added"

	_, err := res.storeBasket.GetBasket(userID)
	if err != nil {
		return domain.Order{}, errors.Wrap(err, errStr)
	}

	basketProducts, err := res.storeBasket.GetBasketProducts(userID)
	if err != nil {
		return domain.Order{}, errors.Wrap(err, errStr)
	}

	if len(basketProducts) == 0 {
		return domain.Order{}, errors.Wrap(domain.ErrBasketEmpty, errStr)
	}

	var productsOrder []domain.ProductOrderForAddOrder

	for _, v := range basketProducts {
		productOrder := domain.ProductOrderForAddOrder{}
		productOrder.ProductID = v.ProductID
		productOrder.Count = v.Count
		productOrder.BasketProductID = v.ID

		product, errGet := res.storeProduct.GetProduct(v.ProductID)
		if errGet != nil {
			return domain.Order{}, errors.Wrap(errGet, errStr)
		}

		productOrder.Price = v.Count * product.Price

		productsOrder = append(productsOrder, productOrder)
	}

	orderID, err := res.store.AddOrder(userID)
	if err != nil {
		return domain.Order{}, errors.Wrap(err, errStr)
	}

	if orderID == 0 {
		return domain.Order{}, errors.Wrap(domain.ErrOrderNotCreated, errStr)
	}

	newOrder, err := res.store.GetOrder(orderID)
	if err != nil {
		return domain.Order{}, errors.Wrap(err, errStr)
	}

	if reflect.DeepEqual(newOrder, domain.Order{}) {
		return domain.Order{}, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	for _, v := range productsOrder {
		productOrder := domain.ProductOrder{}
		productOrder.OrderID = orderID
		productOrder.ProductID = v.ProductID
		productOrder.Count = v.Count
		productOrder.Price = v.Price

		_, err = res.store.AddProductOrder(productOrder)
		if err != nil {
			return domain.Order{}, errors.Wrap(err, errStr)
		}

		isDeleted, errDel := res.storeBasket.DeleteBasketProduct(v.BasketProductID)
		if errDel != nil {
			return domain.Order{}, errors.Wrap(errDel, errStr)
		}
		if !isDeleted {
			return domain.Order{}, errors.Wrap(errDel, errStr)
		}
	}

	return newOrder, nil
}

func (res *OrderService) GetOrder(orderID int) (domain.Order, error) {
	errStr := fmt.Sprintf("order (orderID %d) not fetched", orderID)

	order, err := res.store.GetOrder(orderID)
	if err != nil {
		return domain.Order{}, errors.Wrap(err, errStr)
	}

	if reflect.DeepEqual(order, domain.Order{}) {
		return domain.Order{}, errors.Wrap(domain.ErrUserNotFound, errStr)
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
		return errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	isDeleted, err = res.store.DeleteOrder(orderID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return nil
}

func (res *OrderService) GetOrders() ([]domain.Order, error) {
	errStr := "orders not fetched"
	c, err := res.store.GetOrders()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
