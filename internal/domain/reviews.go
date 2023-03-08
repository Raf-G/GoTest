package domain

type Review struct {
	ID          int    `json:"id" example:"0"`
	UserID      int    `json:"user_id" example:"1"'`
	ProductID   int    `json:"product_id" example:"1"`
	Description string `json:"description" example:"test description"`
	Grade       int    `json:"grade" example:"5"`
}
