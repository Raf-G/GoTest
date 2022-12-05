package service

import (
	"fmt"

	"example.com/m/v2/database/localstorage"
)

func AddCart() {
	cart := localstorage.PostCart()
	fmt.Printf("Cart added %+v \n", cart)
}
