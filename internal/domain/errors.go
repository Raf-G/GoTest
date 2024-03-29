package domain

import "github.com/pkg/errors"

var ErrNotPositiveQuantity = errors.New("quantity must be greater than 0")

// Users
var ErrUserNotFound = errors.New("user not found")
var ErrUserNotEdited = errors.New("user not edited")
var ErrNoLogin = errors.New("login is missing")
var ErrNoSurname = errors.New("surname is missing")

// Roles
var ErrRoleNotFound = errors.New("role not found")

// Basket
var ErrBasketProductNoBasketID = errors.New("basket_id is missing")
var ErrBasketProductNoProductID = errors.New("product_id is missing")
var ErrBasketProductNoCount = errors.New("count is missing")
var ErrBasketProductNotFound = errors.New("basket product not found")
var ErrBasketNotFound = errors.New("basket not found")
var ErrBasketNotDeleted = errors.New("basket not deleted")
var ErrBasketEmpty = errors.New("basket empty")

// Products
var ErrProductNotCreated = errors.New("product not created")
var ErrProductNotFound = errors.New("product not found")

// Reviews
var ErrReviewNotFound = errors.New("review not found")
var ErrReviewNotEdited = errors.New("reviews not edited")

// Orders
var ErrOrderNotCreated = errors.New("order not created")
var ErrOrderNotFound = errors.New("order not found")

// Statuses
var ErrStatusNotFound = errors.New("status not found")
