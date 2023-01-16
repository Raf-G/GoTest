package repository

import (
	"context"
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

func (rep *UserRepository) Add(item domain.User) (*domain.User, error) {
	errStr := "[repository] user not added to the database"

	query := "INSERT INTO `users` (`login`, `name`, `surname`, `password`, `role_id`) VALUES (?, ?, ?, ?, ?)"
	insertResult, err := rep.db.ExecContext(context.Background(), query, item.Login, item.Name, item.Surname, item.Password, item.Role)
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	log.Printf("inserted id: %d", id)

	return &item, nil
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
