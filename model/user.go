package model

import (
	"fmt"
)

type UserDefinition struct {
	Username string `yaml:"Username" json:"USERNAME"`
	Password string `yaml:"Password" json:"PASSWORD"`
	Role     string `yaml:"Role" json:"ROLE"`
}

var Collection map[int]UserDefinition

func (u UserDefinition) Users() map[int]UserDefinition {
	return Collection
}

func (u UserDefinition) Delete(id int) UserDefinition {
	user := Collection[id]
	delete(Collection, id)
	return user
}

func (u UserDefinition) GetUserById(id int) UserDefinition {
	if _, ok := Collection[id]; !ok {
		fmt.Println("not found user")
		return UserDefinition{}
	}
	return u.Users()[id]
}

func (u UserDefinition) Update(id int, user UserDefinition) UserDefinition {
	Collection[id] = user
	return Collection[id]
}

func (u UserDefinition) Create(user UserDefinition) UserDefinition {
	if Collection == nil {
		Collection = map[int]UserDefinition{1: {user.Username, user.Password, user.Role}}
	} else {
		Collection[len(u.Users())+1] = user
	}
	return user
}
