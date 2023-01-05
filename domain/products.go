package domain

type ProductsStorage interface {
	GetProduct(int) (Product, error)
	GetProducts() ([]Product, error)
}

type ProductsService interface {
	GetOneProduct(int) (Product, error)
	GetAllProducts() ([]Product, error)
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
}
