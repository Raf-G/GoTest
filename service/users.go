package service

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/validation"
	"fmt"
	"github.com/pkg/errors"
)

type UserService struct {
	store domain.UsersStorage
}

func NewUserService(storage domain.UsersStorage) *UserService {
	return &UserService{storage}
}

func (is *UserService) Add(item domain.User) (domain.User, error) {
	errStr := fmt.Sprintf("[services] item not added")

	err := validation.UserValidation(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	itemDB, err := is.store.Add(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	if itemDB == nil {
		return item, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return *itemDB, nil
}

func (cs *UserService) GetAll() ([]domain.User, error) {
	errStr := fmt.Sprintf("[services] users not fetched")
	c, err := cs.store.GetUsers()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
