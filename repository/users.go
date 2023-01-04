package repository

import (
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (rep *UserRepository) GetUsers() ([]domain.User, error) {
	rows, err := rep.db.Query("select * from toy_shop.users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	users := []domain.User{}

	for rows.Next() {
		p := domain.User{}
		err := rows.Scan(&p.ID, &p.Login, &p.Surname, &p.Name, &p.Password, &p.Role)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, p)
	}
	return users, nil
}
