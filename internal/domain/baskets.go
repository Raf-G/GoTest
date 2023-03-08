package domain

type BasketProduct struct {
	ID         int `json:"id" example:"0"`
	BasketID   int `json:"basket_id" example:"1"`
	ProductID  int `json:"product_id" example:"1"`
	Count      int `json:"count" example:"2"`
	TotalPrice int `json:"total_price" example:"400"`
}

type Basket struct {
	ID         int             `json:"id" example:"0"`
	UserID     int             `json:"user_id" example:"1"`
	Products   []BasketProduct `json:"products"`
	TotalPrice int             `json:"total_price" example:"400"`
}
