package repository

import (
	"database/sql"
	"example.com/m/v2/internal/domain"
)

//go:generate mockgen -source=roles.go -destination=mocks/roles.go

type RolesStorage interface {
	GetRole(int) (domain.Role, error)
	GetRoles() ([]domain.Role, error)
}

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (res *RoleRepository) GetRole(id int) (domain.Role, error) {
	row := res.db.QueryRow("SELECT * FROM roles WHERE id = ?", id)

	role := domain.Role{}

	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return domain.Role{}, err
	}

	return role, nil
}

func (rep *RoleRepository) GetRoles() ([]domain.Role, error) {
	rows, err := rep.db.Query("select * from toy_shop.roles")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []domain.Role

	for rows.Next() {
		p := domain.Role{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, p)
	}
	return roles, nil
}
