package service

import (
	"fmt"

	"example.com/m/v2/database/localstorage"
	"example.com/m/v2/domain"
)

func AddCartItem() {
	var addItem domain.Item

	fmt.Println("Input your cart id")
	fmt.Scanf("%d\n", &addItem.CartID)

	_, ok := localstorage.GetCart(addItem.CartID)
	if !ok {
		fmt.Println("Cart not found")
		return
	}
	for addItem.Name == "" {
		fmt.Println("Input your product name")
		_, err := fmt.Scanf("%s\n", &addItem.Name)
		if err != nil {
			fmt.Println(err)
		}
	}

	for addItem.Quantity == 0 {
		fmt.Printf("Input quantity for %s\n", addItem.Name)
		_, err := fmt.Scanf("%d\n", &addItem.Quantity)
		if err != nil {
			fmt.Println(err)
		}
	}

	value, _ := localstorage.PostItem(addItem.CartID, addItem)

	fmt.Printf("Item added to cart %+v\n", value)
	addItem = domain.Item{}
}

func DeleteCartItem() {
	var inputCartID domain.CartID
	var inputItemID domain.ItemID
	for inputCartID == 0 {
		fmt.Println("Input cart id")
		fmt.Scanf("%d\n", &inputCartID)
	}

	for inputItemID == 0 {
		fmt.Println("Input item id")
		fmt.Scanf("%d\n", &inputItemID)
	}

	value, _ := localstorage.DeleteItem(inputCartID, inputItemID)

	fmt.Printf("Item removed from cart %+v\n", value)
}

func printMessage(message localstorage.MessageShowCartItems) {
	message.ShowCartItemsMessage()
}

func ShowCartItems() {
	var inputCartID domain.CartID
	for inputCartID == 0 {
		var text localstorage.ItemsText = "Input cart id"
		printMessage(text)
		fmt.Scanf("%d\n", &inputCartID)
	}

	value, _ := localstorage.GetCartItems(inputCartID)
	var msg = value
	printMessage(msg)
}
