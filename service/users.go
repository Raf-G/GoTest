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

func (res *UserService) Add(item domain.User) (domain.User, error) {
	errStr := "[services] item not added"

	err := validation.UserValidation(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	itemDB, err := res.store.Add(item)
	if err != nil {
		return item, errors.Wrap(err, errStr)
	}

	if itemDB == nil {
		return item, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return *itemDB, nil
}

func (res *UserService) GetUser(id int) (domain.User, error) {
	errStr := fmt.Sprintf("[services] user (userID %d) not fetched", id)

	user, err := res.store.GetUser(id)
	if err != nil {
		return domain.User{}, errors.Wrap(err, errStr)
	}

	if user == nil {
		return domain.User{}, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return *user, err
}

func (res *UserService) GetAll() ([]domain.User, error) {
	errStr := "[services] users not fetched"
	c, err := res.store.GetUsers()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}

func (res *UserService) Edit(user domain.User) (*domain.User, error) {
	errStr := "[services] user not edit"

	newUser, err := res.store.Edit(user)
	if err != nil {
		return nil, errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return newUser, nil
}

func (res *UserService) Delete(userID int) error {
	errStr := fmt.Sprintf("[services] user (userID %d) not deleted", userID)

	isDeleted, err := res.store.Delete(userID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain.ErrUserNotFound, errStr)
	}

	return nil
}
