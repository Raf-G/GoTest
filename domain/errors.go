package domain

import "github.com/pkg/errors"

var ErrItemNotFound = errors.New("item not found")

var ErrUserNotFound = errors.New("user not found")

var ErrNoLogin = errors.New("login is missing")

var ErrNoSurname = errors.New("surname is missing")

var ErrNotPositiveQuantity = errors.New("quantity must be greater than 0")
