package service

import (
	"example.com/m/v2/database/localstorage"
	"fmt"
)

func AddCart() {
	cart := localstorage.PostCart()

	fmt.Printf("Cart added %+v \n", cart)
}
