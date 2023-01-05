package domain

type BasketsStorage interface {
	GetBasket(int) (Basket, error)
}

type BasketsService interface {
	GetBasket(int) (Basket, error)
}

type BasketProduct struct {
	ProductID  int
	Count      int
	TotalPrice int
}

type Basket struct {
	ID         int
	UserID     int
	Products   []BasketProduct
	TotalPrice int
}
