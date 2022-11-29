package domain

type CartID int

type Cart struct {
	Id    CartID
	items []Item
}
