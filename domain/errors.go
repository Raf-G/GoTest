package domain

import "github.com/pkg/errors"

var ErrItemNotFound = errors.New("item not found")
var ErrNotPositiveQuantity = errors.New("quantity must be greater than 0")

//Users
var ErrUserNotFound = errors.New("user not found")
var ErrNoLogin = errors.New("login is missing")
var ErrNoSurname = errors.New("surname is missing")

//Roles
var ErrRoleNotFound = errors.New("role not found")

//Basket
var ErrBasketProductNoBasketID = errors.New("basket_id is missing")
var ErrBasketProductNoProductID = errors.New("product_id is missing")
var ErrBasketProductNoCount = errors.New("count is missing")
var ErrBasketProductNotFound = errors.New("basket product not found")

//Products
var ErrProductNotCreated = errors.New("product not created")
var ErrProductNotFound = errors.New("product not found")
