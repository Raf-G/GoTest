package domain

type UsersStorage interface {
	Add(User) (*User, error)
	GetUsers() ([]User, error)
}

type UsersService interface {
	Add(User) (User, error)
	GetAll() ([]User, error)
}

type User struct {
	ID       int
	Login    string
	Surname  string
	Name     string
	Password string
	Role     int
}
