package cartStruct

type CartID int
type ItemID int
type ItemByIDMap map[ItemID]Item

type Item struct {
	Id       ItemID
	CartID   CartID
	Name     string
	Quantity int
}

type Cart struct {
	Id    CartID
	items []Item
}
