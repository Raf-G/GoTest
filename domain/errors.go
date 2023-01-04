package domain

import "github.com/pkg/errors"

var ErrItemNotFound = errors.New("item not found")

var ErrCartNotFound = errors.New("cart not found")

var ErrNoProduct = errors.New("product is missing")

var ErrNotPositiveQuantity = errors.New("quantity must be greater than 0")
