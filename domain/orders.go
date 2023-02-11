package domain

//go:generate mockgen -source=users.go -destination=mocks/mock.go

type OrdersStorage interface {
	AddOrder(int) (int, error)
	GetOrder(int) (Order, error)
	DeleteOrder(int) (bool, error)
	GetOrders() ([]Order, error)
	AddProductOrder(ProductOrder) (int, error)
	DeleteProductsOrder(int) (bool, error)
}

type OrdersService interface {
	AddOrder(int) (Order, error)
	GetOrder(int) (Order, error)
	DeleteOrder(int) error
	GetOrders() ([]Order, error)
}

type Order struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	StatusID int `json:"status_id"`
}

type ProductOrder struct {
	ID        int
	OrderID   int
	ProductID int
	Count     int
	Price     int
}

type ProductOrderForAddOrder struct {
	ProductOrder
	BasketProductID int
}
