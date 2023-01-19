package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (res *OrderRepository) AddOrder(userID int) (int, error) {
	errStr := "[repository] order not added to the database"

	query := "INSERT INTO `orders` (`user_id`) VALUES (?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, userID)
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return 0, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return 0, err
	}

	log.Printf("inserted id: %d", id)

	return int(id), nil
}

func (res *OrderRepository) GetOrder(orderID int) (domain.Order, error) {
	errStr := "[repository] order not fetched from the database: "

	row := res.db.QueryRow("SELECT * FROM orders WHERE id = ?", orderID)

	order := domain.Order{}

	err := row.Scan(&order.ID, &order.UserID, &order.StatusID)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.Order{}, err
	}

	return order, nil
}

func (res *OrderRepository) DeleteOrder(orderID int) (bool, error) {
	errStr := "[repository] order not deleted from the database: "

	_, err := res.db.Exec("DELETE FROM orders WHERE id = ?", orderID)
	if err != nil {
		fmt.Println(errStr, err)
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
			fmt.Println(err)
			return nil, err
		}
		orders = append(orders, p)
	}
	return orders, nil
}

func (res *OrderRepository) AddProductOrder(productOrder domain.ProductOrder) (int, error) {
	errStr := "[repository] product order not added to the database"

	query := "INSERT INTO `products_orders` (`order_id`, `product_id`, `count`, `price`) VALUES (?,?,?,?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, productOrder.OrderID, productOrder.ProductID, productOrder.Count, productOrder.Price)
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return 0, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Printf("%s: %s", errStr, err)
		return 0, err
	}

	log.Printf("inserted id: %d", id)

	return int(id), nil
}

func (res *OrderRepository) DeleteProductsOrder(orderID int) (bool, error) {
	errStr := "[repository] product order not deleted from the database: "

	_, err := res.db.Exec("DELETE FROM products_orders WHERE order_id = ?", orderID)
	if err != nil {
		fmt.Println(errStr, err)
		return false, err
	}

	return true, nil
}
