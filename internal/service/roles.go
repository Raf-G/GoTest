package service

import (
	"example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repository"
	"fmt"
	"github.com/pkg/errors"
)

type RolesService interface {
	GetRole(int) (domain.Role, error)
	GetRoleAll() ([]domain.Role, error)
}

type RoleService struct {
	store repository.RolesStorage
}

func NewRoleService(storage repository.RolesStorage) *RoleService {
	return &RoleService{storage}
}

func (res *RoleService) GetRole(id int) (domain.Role, error) {
	errStr := fmt.Sprintf("role (roleID %d) not fetched", id)

	role, err := res.store.GetRole(id)
	if err != nil {
		return domain.Role{}, errors.Wrap(err, errStr)
	}
	if role == nil {
		return domain.Role{}, errors.Wrap(domain.ErrRoleNotFound, errStr)
	}

	return *role, err
}

func (cs *RoleService) GetRoleAll() ([]domain.Role, error) {
	errStr := "users not fetched"
	c, err := cs.store.GetRoles()
	if err != nil {
		return nil, errors.Wrap(err, errStr)
	}

	return c, nil
}
