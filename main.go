package main

import (
	"fmt"

	"example.com/m/v2/service"
)

func main() {
	var inputNumber int

	for {
		fmt.Println("1. Add cart\n2. Add cart item to cart\n3. Remove item from cart\n4. View cart")
		fmt.Scanf("%d\n", &inputNumber)

		switch inputNumber {
		case 1:
			service.AddCart()
		case 2:
			service.AddCartItem()
		case 3:
			service.DeleteCartItem()
		case 4:
			service.ShowCartItems()
		case 5:
			service.ShowCartItems()
		default:
			fmt.Println("This case is not exist")
		}
	}
}
