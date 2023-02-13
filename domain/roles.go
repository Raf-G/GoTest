package domain

//go:generate mockgen -source=roles.go -destination=mocks/roles.go

type RolesStorage interface {
	GetRole(int) (Role, error)
	GetRoles() ([]Role, error)
}

type RolesService interface {
	GetRole(int) (Role, error)
	GetRoleAll() ([]Role, error)
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
