package domain

type Product struct {
	ID          int    `json:"id" example:"0"`
	Name        string `json:"name" example:"testProduct"`
	Description string `json:"description" example:"test description"`
	Price       int    `json:"price" example:"300"`
}
