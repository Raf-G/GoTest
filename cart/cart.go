package cart

import (
	"example.com/m/v2/database"
	"fmt"
)

func AddCart() {
	database.PutMaxIDCardMapIncrement()
	database.PostCart(database.MaxIDCardMap)

	value, _ := database.GetCart(database.MaxIDCardMap)

	fmt.Printf("Cart added %+v \n", value)
}
