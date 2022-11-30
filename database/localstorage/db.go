package localstorage

import (
	"example.com/m/v2/domain"
	"fmt"
)

type ItemByIDMap map[domain.ItemID]domain.Item

type ItemsText string

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

type MessageShowCartItems interface {
	ShowCartItemsMessage()
}

func (text ItemsText) ShowCartItemsMessage() {
	fmt.Println(text)
}

func (value ItemByIDMap) ShowCartItemsMessage() {
	fmt.Printf("Your cart %+v\n", value)
}
