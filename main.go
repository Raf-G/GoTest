package main

import (
	"example.com/m/v2/cart"
	"example.com/m/v2/cartItem"
	"fmt"
)

func main() {
	var inputNumber int

	for {
		fmt.Println("1. Add cart\n2. Add cart item to cart\n3. Remove item from cart\n4. View cart")
		fmt.Scanf("%d\n", &inputNumber)

		switch inputNumber {
		case 1:
			cart.AddCart()
		case 2:
			cartItem.AddCartItem()
		case 3:
			cartItem.DeleteCartItem()
		case 4:
			cartItem.ShowCartItems()
		default:
			fmt.Println("This case is not exist")
		}
	}
}
