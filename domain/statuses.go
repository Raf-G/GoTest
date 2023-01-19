package domain

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
