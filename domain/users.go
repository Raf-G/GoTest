package domain

//go:generate mockgen -source=users.go -destination=mocks/mock.go

type UsersStorage interface {
	Add(User) (User, error)
	GetUser(int) (*User, error)
	GetUsers() ([]User, error)
	Edit(User) (User, error)
	Delete(int) (bool, error)
}

type UsersService interface {
	Add(User) (User, error)
	GetUser(int) (User, error)
	GetAll() ([]User, error)
	Edit(User) (User, error)
	Delete(int) error
}

type User struct {
	ID       int
	Login    string
	Surname  string
	Name     string
	Password string
	Role     int
}
