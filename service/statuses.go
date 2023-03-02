package service

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/repository"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

//go:generate mockgen -source=statuses.go -destination=mocks/statuses.go

type StatusesService interface {
	GetStatus(int) (domain.Status, error)
	GetStatuses() ([]domain.Status, error)
}

type StatusService struct {
	store repository.StatusesStorage
}

func NewStatusService(storage repository.StatusesStorage) *StatusService {
	return &StatusService{storage}
}

func (res *StatusService) GetStatus(statusID int) (domain.Status, error) {
	errStr := fmt.Sprintf("role (roleID %d) not fetched", statusID)

	status, err := res.store.GetStatus(statusID)
	if err != nil {
		return domain.Status{}, errors.Wrap(err, errStr)
	}

	if reflect.DeepEqual(status, domain.Status{}) {
		return domain.Status{}, errors.Wrap(domain.ErrStatusNotFound, errStr)
	}

	return status, err
}

func (res *StatusService) GetStatuses() ([]domain.Status, error) {
	errStr := "statuses not fetched"
	c, err := res.store.GetStatuses()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
