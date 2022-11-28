package database

import cartStruct "example.com/m/v2/struct"

var CartMap = make(map[cartStruct.CartID]cartStruct.Cart)
var ItemMap = make(map[cartStruct.CartID]cartStruct.ItemByIDMap)
var MaxIDCardMap cartStruct.CartID = 0
var MaxIDItemMap cartStruct.ItemID = 0

func PostCart(maxIDCardMap cartStruct.CartID) {
	CartMap[maxIDCardMap] = cartStruct.Cart{Id: maxIDCardMap}
}

func GetCart(cartID cartStruct.CartID) (cartStruct.Cart, bool) {
	value, ok := CartMap[cartID]
	return value, ok
}

func PostItem(cartID cartStruct.CartID, newItem cartStruct.Item) {
	_, ok := GetCartItems(cartID)
	if !ok {
		item := make(map[cartStruct.ItemID]cartStruct.Item)
		item[MaxIDItemMap] = newItem
		ItemMap[cartID] = item
	} else {
		ItemMap[cartID][MaxIDItemMap] = newItem
	}
}

1func GetCartItems(cartID cartStruct.CartID) (cartStruct.ItemByIDMap, bool) {
	value, ok := ItemMap[cartID]
	return value, ok
}

func DeleteItem(cartID cartStruct.CartID, itemID cartStruct.ItemID) {
	delete(ItemMap[cartID], itemID)
}

func PutMaxIDCardMapIncrement() {
	MaxIDCardMap++
}

func PutMaxIDItemMapIncrement() {
	MaxIDItemMap++
}

func GetMaxIDItemMapIncrement() cartStruct.ItemID {
	return MaxIDItemMap
}
