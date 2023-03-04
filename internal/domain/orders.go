package domain

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
