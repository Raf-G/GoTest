package domain

type ItemID int

type Item struct {
	Id       ItemID
	CartID   CartID
	Name     string
	Quantity int
}
