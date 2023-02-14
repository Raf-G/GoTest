package domain

//go:generate mockgen -source=statuses.go -destination=mocks/statuses.go

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
