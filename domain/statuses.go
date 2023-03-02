package domain

//go:generate mockgen -source=statuses.go -destination=mocks/statuses.go

type Status struct {
	ID   int
	Name string
}
