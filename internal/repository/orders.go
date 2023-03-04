package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/internal/domain"
	"log"
)

//go:generate mockgen -source=orders.go -destination=mocks/orders.go

type OrdersStorage interface {
	AddOrder(int) (int, error)
	GetOrder(int) (domain.Order, error)
	DeleteOrder(int) (bool, error)
	GetOrders() ([]domain.Order, error)
	AddProductOrder(domain.ProductOrder) (int, error)
	DeleteProductsOrder(int) (bool, error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (res *OrderRepository) AddOrder(userID int) (int, error) {
	query := "INSERT INTO `orders` (`user_id`) VALUES (?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, userID)
	if err != nil {
		return 0, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (res *OrderRepository) GetOrder(orderID int) (domain.Order, error) {
	row := res.db.QueryRow("SELECT * FROM orders WHERE id = ?", orderID)

	order := domain.Order{}

	err := row.Scan(&order.ID, &order.UserID, &order.StatusID)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (res *OrderRepository) DeleteOrder(orderID int) (bool, error) {
	_, err := res.db.Exec("DELETE FROM orders WHERE id = ?", orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (res *OrderRepository) GetOrders() ([]domain.Order, error) {
	rows, err := res.db.Query("select * from orders")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		p := domain.Order{}
		err = rows.Scan(&p.ID, &p.UserID, &p.StatusID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, p)
	}
	return orders, nil
}

func (res *OrderRepository) AddProductOrder(productOrder domain.ProductOrder) (int, error) {
	query := "INSERT INTO `products_orders` (`order_id`, `product_id`, `count`, `price`) VALUES (?,?,?,?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, productOrder.OrderID, productOrder.ProductID, productOrder.Count, productOrder.Price)
	if err != nil {
		return 0, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (res *OrderRepository) DeleteProductsOrder(orderID int) (bool, error) {
	_, err := res.db.Exec("DELETE FROM products_orders WHERE order_id = ?", orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}
