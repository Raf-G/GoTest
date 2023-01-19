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

func (res *UserRepository) Add(item domain.User) (*domain.User, error) {
	errStr := "[repository] user not added to the database"

	query := "INSERT INTO `users` (`login`, `name`, `surname`, `password`, `role_id`) VALUES (?, ?, ?, ?, ?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, item.Login, item.Name, item.Surname, item.Password, item.Role)
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("%s: %s", errStr, err)
	}
	item.ID = int(id)
	log.Printf("inserted id: %d", id)

	return &item, nil
}

func (res *UserRepository) GetUser(id int) (domain.User, error) {
	errStr := "[repository] user not fetched from the database: "

	row := res.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := domain.User{}

	err := row.Scan(&user.ID, &user.Login, &user.Surname, &user.Name, &user.Password, &user.Role)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.User{}, err
	}

	return user, nil
}

func (res *UserRepository) GetUsers() ([]domain.User, error) {
	rows, err := res.db.Query("select * from toy_shop.users")
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

func (res *UserRepository) Edit(user domain.User) (domain.User, error) {
	errStr := "[repository] user not edit from the database: "

	stmt, err := res.db.Prepare("UPDATE users SET login = ?, name = ?, surname = ? , password = ?, role_id = ? WHERE id = ?")
	if err != nil {
		fmt.Println(errStr, err)
		return domain.User{}, err
	}

	_, err = stmt.Exec(user.Login, user.Name, user.Surname, user.Password, user.Role, user.ID)
	if err != nil {
		fmt.Println(errStr, err)
		return domain.User{}, err
	}

	return user, nil
}

func (res *UserRepository) Delete(userID string) (bool, error) {
	errStr := "[repository] user not deleted from the database: "

	_, err := res.db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		fmt.Println(errStr, err)
		return false, err
	}

	return true, nil
}
