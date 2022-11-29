package localstorage

import "example.com/m/v2/domain"

type ItemByIDMap map[domain.ItemID]domain.Item

var CartMap = make(map[domain.CartID]domain.Cart)
var ItemMap = make(map[domain.CartID]ItemByIDMap)
var MaxIDCardMap domain.CartID = 0
var MaxIDItemMap domain.ItemID = 0

func putMaxIDCardMapIncrement() {
	MaxIDCardMap++
}

func PutMaxIDItemMapIncrement() domain.ItemID {
	MaxIDItemMap++
	return MaxIDItemMap
}

func GetMaxIDItemMapIncrement() domain.ItemID {
	return MaxIDItemMap
}
