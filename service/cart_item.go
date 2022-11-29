package service

import (
	"example.com/m/v2/database/localstorage"
	cartStruct "example.com/m/v2/domain"
	"fmt"
)

func AddCartItem() {
	var addItem cartStruct.Item

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

	value, ok := localstorage.PostItem(addItem.CartID, addItem)

	fmt.Printf("Item added to cart %+v\n", value)
	addItem = cartStruct.Item{}
}

func DeleteCartItem() {
	var inputCartID cartStruct.CartID
	var inputItemID cartStruct.ItemID
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

func ShowCartItems() {
	var inputCartID cartStruct.CartID
	for inputCartID == 0 {
		fmt.Println("Input cart id")
		fmt.Scanf("%d\n", &inputCartID)
	}

	value, _ := localstorage.GetCartItems(inputCartID)

	fmt.Printf("Your cart %+v\n", value)
}
