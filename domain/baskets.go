package domain

//go:generate mockgen -source=baskets.go -destination=mocks/baskets.go

type BasketProduct struct {
	ID         int `json:"id"`
	BasketID   int `json:"basket_id"`
	ProductID  int `json:"product_id"`
	Count      int `json:"count"`
	TotalPrice int `json:"total_price"`
}

type Basket struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	Products   []BasketProduct
	TotalPrice int `json:"total_price"`
}
