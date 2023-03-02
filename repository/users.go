package repository

import (
	"context"
	"database/sql"
	"example.com/m/v2/domain"
)

//go:generate mockgen -source=users.go -destination=mocks/users.go

type UsersStorage interface {
	Add(domain.User) (domain.User, error)
	GetUser(int) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	Edit(domain.User) (domain.User, error)
	Delete(int) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (res *UserRepository) Add(u domain.User) (domain.User, error) {
	query := "INSERT INTO `users` (`login`, `name`, `surname`, `password`, `role_id`) VALUES (?, ?, ?, ?, ?)"
	insertResult, err := res.db.ExecContext(context.Background(), query, u.Login, u.Name, u.Surname, u.Password, u.Role)
	if err != nil {
		return u, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = int(id)
	return u, nil
}

func (res *UserRepository) GetUser(id int) (*domain.User, error) {
	row := res.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := domain.User{}

	err := row.Scan(&user.ID, &user.Login, &user.Surname, &user.Name, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (res *UserRepository) GetUsers() ([]domain.User, error) {
	rows, err := res.db.Query("select * from toy_shop.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []domain.User

	for rows.Next() {
		p := domain.User{}
		err = rows.Scan(&p.ID, &p.Login, &p.Surname, &p.Name, &p.Password, &p.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, p)
	}
	return users, nil
}

func (res *UserRepository) Edit(user domain.User) (domain.User, error) {
	stmt, err := res.db.Prepare("UPDATE users SET login = ?, name = ?, surname = ? , password = ?, role_id = ? WHERE id = ?")
	if err != nil {
		return domain.User{}, err
	}

	_, err = stmt.Exec(user.Login, user.Name, user.Surname, user.Password, user.Role, user.ID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (res *UserRepository) Delete(userID int) (bool, error) {
	_, err := res.db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
