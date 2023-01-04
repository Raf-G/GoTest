package domain

type UsersStorage interface {
	GetUsers() ([]User, error)
}

type UsersService interface {
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
