package handlers

import "example.com/m/v2/domain"

type jsonUser struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func jsonUserFromUser(item domain.User) jsonUser {
	return jsonUser{ID: item.ID, Login: item.Login, Surname: item.Surname, Name: item.Name, Password: item.Password, Role: item.Role}
}

func userFromJSONUser(jsonItem jsonUser) domain.User {
	return domain.User{ID: jsonItem.ID, Login: jsonItem.Login, Surname: jsonItem.Surname, Name: jsonItem.Name, Password: jsonItem.Password, Role: jsonItem.Role}
}
