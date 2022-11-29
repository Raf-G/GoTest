package localstorage

import "example.com/m/v2/domain"

func PostItem(cartID domain.CartID, newItem domain.Item) (ItemByIDMap, bool) {
	maxIDItemMap := PutMaxIDItemMapIncrement()
	newItem.Id = maxIDItemMap
	_, ok := GetCartItems(cartID)
	if !ok {
		item := make(map[domain.ItemID]domain.Item)
		item[maxIDItemMap] = newItem
		ItemMap[cartID] = item
	} else {
		ItemMap[cartID][maxIDItemMap] = newItem
	}
	value, ok := GetCartItems(cartID)
	return value, ok
}

func GetCartItems(cartID domain.CartID) (ItemByIDMap, bool) {
	value, ok := ItemMap[cartID]
	return value, ok
}

func DeleteItem(cartID domain.CartID, itemID domain.ItemID) (ItemByIDMap, bool) {
	delete(ItemMap[cartID], itemID)
	value, ok := GetCartItems(cartID)
	return value, ok
}
