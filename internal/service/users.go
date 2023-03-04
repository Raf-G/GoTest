package service

import (
	domain2 "example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/validation"
	"fmt"
	"github.com/pkg/errors"
)

type UsersService interface {
	Add(domain2.User) (domain2.User, error)
	GetUser(int) (domain2.User, error)
	GetAll() ([]domain2.User, error)
	Edit(domain2.User) (domain2.User, error)
	Delete(int) error
}

type UserService struct {
	store repository.UsersStorage
}

func NewUserService(storage repository.UsersStorage) *UserService {
	return &UserService{storage}
}

func (res *UserService) Add(u domain2.User) (domain2.User, error) {
	errStr := "user not added"

	err := validation.UserValidation(u)
	if err != nil {
		return u, errors.Wrap(err, errStr)
	}

	userDB, err := res.store.Add(u)
	if err != nil {
		return u, errors.Wrap(err, errStr)
	}

	return userDB, nil
}

func (res *UserService) GetUser(id int) (domain2.User, error) {
	errStr := fmt.Sprintf("user (userID %d) not fetched", id)

	user, err := res.store.GetUser(id)
	if err != nil {
		return domain2.User{}, errors.Wrap(err, errStr)
	}

	if user == nil {
		return domain2.User{}, errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	return *user, err
}

func (res *UserService) GetAll() ([]domain2.User, error) {
	errStr := "users not fetched"
	c, err := res.store.GetUsers()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}

func (res *UserService) Edit(u domain2.User) (domain2.User, error) {
	errStr := "user not edit"

	err := validation.UserValidation(u)
	if err != nil {
		return u, errors.Wrap(err, errStr)
	}

	newUser, err := res.store.Edit(u)
	if err != nil {
		return domain2.User{}, errors.Wrap(domain2.ErrUserNotEdited, errStr)
	}

	return newUser, nil
}

func (res *UserService) Delete(userID int) error {
	errStr := fmt.Sprintf("user (userID %d) not deleted", userID)

	isDeleted, err := res.store.Delete(userID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	if !isDeleted {
		return errors.Wrap(domain2.ErrUserNotFound, errStr)
	}

	return nil
}
