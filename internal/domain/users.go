package domain

type User struct {
	ID       int    `json:"id" example:"0"`
	Login    string `json:"login" example:"testLogin"`
	Surname  string `json:"surname" example:"surnameTest"`
	Name     string `json:"name" example:"nameTest"`
	Password string `json:"password" example:"qweqwe122"`
	Role     int    `json:"role" example:"1"`
}
