package handlers

import (
	domain2 "example.com/m/v2/internal/domain"
)

type jsonUser struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type jsonRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func jsonUserFromUser(u domain2.User) jsonUser {
	return jsonUser{ID: u.ID, Login: u.Login, Surname: u.Surname, Name: u.Name, Password: u.Password, Role: u.Role}
}

func userFromJSONUser(u jsonUser) domain2.User {
	return domain2.User{ID: u.ID, Login: u.Login, Surname: u.Surname, Name: u.Name, Password: u.Password, Role: u.Role}
}

func jsonRoleFromRole(u domain2.Role) jsonRole {
	return jsonRole{ID: u.ID, Name: u.Name}
}
