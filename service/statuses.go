package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

type StatusService struct {
	store domain.StatusesStorage
}

func NewStatusService(storage domain.StatusesStorage) *StatusService {
	return &StatusService{storage}
}

func (res *StatusService) GetStatus(statusID int) (domain.Status, error) {
	errStr := fmt.Sprintf("[services] role (roleID %d) not fetched", statusID)

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
	errStr := "[services] statuses not fetched"
	c, err := res.store.GetStatuses()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
