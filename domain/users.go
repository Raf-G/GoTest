package domain

//go:generate mockgen -source=users.go -destination=mocks/users.go

type User struct {
	ID       int
	Login    string
	Surname  string
	Name     string
	Password string
	Role     int
}
