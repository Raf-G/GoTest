package domain

//go:generate mockgen -source=users.go -destination=mocks/mock.go

type StatusesStorage interface {
	GetStatus(int) (Status, error)
	GetStatuses() ([]Status, error)
}

type StatusesService interface {
	GetStatus(int) (Status, error)
	GetStatuses() ([]Status, error)
}

type Status struct {
	ID   int
	Name string
}
