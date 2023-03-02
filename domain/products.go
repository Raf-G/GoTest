package domain

//go:generate mockgen -source=products.go -destination=mocks/products.go

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
