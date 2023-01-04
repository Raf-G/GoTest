package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
)

type UserService struct {
	store domain.UsersStorage
}

func NewUserService(storage domain.UsersStorage) *UserService {
	return &UserService{storage}
}

func (cs *UserService) GetAll() ([]domain.User, error) {
	errStr := fmt.Sprintf("[services] users not fetched")
	c, err := cs.store.GetUsers()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
