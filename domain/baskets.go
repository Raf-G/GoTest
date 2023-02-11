package domain

//go:generate mockgen -source=users.go -destination=mocks/mock.go

type BasketsStorage interface {
	AddBasketProduct(BasketProduct) (BasketProduct, error)
	GetBasketProduct(int, int) (BasketProduct, error)
	GetBasketProducts(int) ([]BasketProduct, error)
	EditBasketProduct(BasketProduct) (BasketProduct, error)
	DeleteBasketProduct(int) (bool, error)
	GetBasket(int) (Basket, error)
}

type BasketsService interface {
	AddProductToBasket(BasketProduct) (BasketProduct, error)
	DecreaseQuantityProductToBasket(BasketProduct) (BasketProduct, error)
	DeleteProductToBasket(int) error
	GetBasket(int) (Basket, error)
}

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
