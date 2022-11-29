package localstorage

import "example.com/m/v2/domain"

func PostCart() domain.Cart {
	putMaxIDCardMapIncrement()
	newCart := domain.Cart{Id: MaxIDCardMap}
	CartMap[MaxIDCardMap] = newCart
	return newCart
}

func GetCart(cartID domain.CartID) (domain.Cart, bool) {
	value, ok := CartMap[cartID]
	return value, ok
}
