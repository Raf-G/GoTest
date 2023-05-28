package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/internal/domain"
	"fmt"
	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=roles.go -destination=mocks/roles.go

type RolesStorage interface {
	GetRole(int) (*domain.Role, error)
	GetRoles() ([]domain.Role, error)
}

type RoleRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewRoleRepository(db *sql.DB, redisClient *redis.Client) *RoleRepository {
	return &RoleRepository{db, redisClient}
}

func (res *RoleRepository) GetRole(id int) (*domain.Role, error) {
	row := res.db.QueryRow("SELECT * FROM roles WHERE id = ?", id)

	role := domain.Role{}

	err := row.Scan(&role.ID, &role.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	err = res.redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := res.redisClient.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("dont find")
	}
	fmt.Println("key", val)

	return &role, nil
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
