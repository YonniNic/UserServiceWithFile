package service

import (
	"UserService/model"
	"UserService/repository"
	"fmt"
)

func NewUserPanel(definition model.UserDefinition, database repository.DataBase) userPanel {
	return userPanel{
		definition,
		database,
	}
}

type userPanel struct {
	model.UserDefinition
	repository.DataBase
}

func (u userPanel) Logging(username, password string) (int, error) {
	for id, user := range u.Users() {
		if username == user.Username && password == user.Password {
			return id, nil
		}
	}
	return 0, fmt.Errorf("not found user")
}

func (u userPanel) ViewDataForUser(id int) model.UserDefinition {
	if u.Role == "Admin" {
		return u.GetUserById(id)
	}
	return model.UserDefinition{}
}

func (u userPanel) DeleteUser(id int) {
	if u.Role == "Admin" {
		u.Delete(id)
	}
}

/*
func (u userPanel) Signup(username, password, role string) error {
	newUser := model.UserDefinition{Username: username, Password: password, Role: role}
	for _, user := range u.AllUsers() {
		if user.Username == newUser.Username {
			return fmt.Errorf("that username is already taken")
		}
	}
	u.CreateUser(newUser)
	//u.Users()[string(len(u.Users())+1)] = newUser
	return fmt.Errorf("successful registration")
}

*/
