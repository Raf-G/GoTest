package domain

//go:generate mockgen -source=users.go -destination=mocks/mock.go

type ProductsStorage interface {
	AddProduct(Product) (*Product, error)
	GetProduct(int) (Product, error)
	EditProduct(Product) (Product, error)
	DeleteProduct(int) (bool, error)
	GetProducts() ([]Product, error)
}

type ProductsService interface {
	AddProduct(Product) (Product, error)
	GetProduct(int) (Product, error)
	EditProduct(Product) (Product, error)
	DeleteProduct(int) error
	GetAllProducts() ([]Product, error)
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
