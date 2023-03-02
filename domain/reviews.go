package domain

//go:generate mockgen -source=reviews.go -destination=mocks/reviews.go

type Review struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	Description string `json:"description"`
	Grade       int    `json:"grade"`
}
