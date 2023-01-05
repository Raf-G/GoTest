package domain

type RolesStorage interface {
	GetRoles() ([]Role, error)
}

type RolesService interface {
	GetRoleAll() ([]Role, error)
}

type Role struct {
	ID   int
	Name string
}
