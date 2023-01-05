package service

import (
	"example.com/m/v2/domain"
	"fmt"
	"github.com/pkg/errors"
)

type RoleService struct {
	store domain.RolesStorage
}

func NewRoleService(storage domain.RolesStorage) *RoleService {
	return &RoleService{storage}
}

func (cs *RoleService) GetRoleAll() ([]domain.Role, error) {
	errStr := fmt.Sprintf("[services] users not fetched")
	c, err := cs.store.GetRoles()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
