package domain

type UsersStorage interface {
	Add(User) (*User, error)
	GetUser(int) (User, error)
	GetUsers() ([]User, error)
	Edit(User) (User, error)
	Delete(string) (bool, error)
}

type UsersService interface {
	Add(User) (User, error)
	GetUser(int) (User, error)
	GetAll() ([]User, error)
	Edit(User) (User, error)
	Delete(string) error
}

type User struct {
	ID       int
	Login    string
	Surname  string
	Name     string
	Password string
	Role     int
}
