package main

import (
	"fmt"
)

type item struct {
	id       itemID
	cartID   cartID
	name     string
	quantity int
}

type cart struct {
	id    cartID
	items []item
}

type cartID int
type itemID int
type itemByIDMap map[itemID]item

func main() {
	var inputNumber int
	cartMap := make(map[cartID]cart)
	itemMap := make(map[cartID]itemByIDMap)
	var maxIDCardMap cartID = 0
	var maxIDItemMap itemID = 0

	var addItem item

	for {
		fmt.Println("1. Add cart\n2. Add cart item to cart\n3. Remove item from cart\n4. View cart")
		fmt.Scanf("%d\n", &inputNumber)

		switch inputNumber {
		case 1:
			maxIDCardMap++
			cartMap[maxIDCardMap] = cart{id: maxIDCardMap}
			fmt.Printf("Cart added %+v \n", cartMap[maxIDCardMap])
		case 2:
			fmt.Println("Input your cart id")
			fmt.Scanf("%d\n", &addItem.cartID)

			_, ok := cartMap[addItem.cartID]
			if !ok {
				fmt.Println("Cart not found")
				break
			}
			for addItem.name == "" {
				fmt.Println("Input your product name")
				_, err := fmt.Scanf("%s\n", &addItem.name)
				if err != nil {
					fmt.Println(err)
				}
			}

			for addItem.quantity == 0 {
				fmt.Printf("Input quantity for %s\n", addItem.name)
				_, err := fmt.Scanf("%d\n", &addItem.quantity)
				if err != nil {
					fmt.Println(err)
				}
			}

			maxIDItemMap++
			addItem.id = maxIDItemMap

			_, ok = itemMap[addItem.cartID]
			if !ok {
				item := make(map[itemID]item)
				item[maxIDItemMap] = addItem
				itemMap[addItem.cartID] = item
			} else {
				itemMap[addItem.cartID][maxIDItemMap] = addItem
			}
			fmt.Printf("Item added to cart %+v\n", itemMap[addItem.cartID])
			addItem = item{}
		case 3:
			var inputCartID int
			var inputItemID int
			for inputCartID == 0 {
				fmt.Println("Input cart id")
				fmt.Scanf("%d\n", &inputCartID)
			}

			for inputItemID == 0 {
				fmt.Println("Input item id")
				fmt.Scanf("%d\n", &inputItemID)
			}

			delete(itemMap[cartID(inputCartID)], itemID(inputItemID))
			fmt.Println("Item removed from cart")
		case 4:
			fmt.Println("Your cart")
		default:
			fmt.Println("This case is not exist")
		}
	}

}
