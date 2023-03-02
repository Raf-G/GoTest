package domain

//go:generate mockgen -source=roles.go -destination=mocks/roles.go

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
