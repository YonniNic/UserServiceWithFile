package repository

import (
	"UserService/model"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

var FileDatabasePath, _ = filepath.Abs("/Users/I556357/Downloads/resources/database.yml")

type FileDatabase struct {
	users map[int]model.UserDefinition `yaml:"Users"`
}

/*
	func (f FileDatabase) GetUserByID(id string) model.UserDefinition {
		if _, ok := f.users[id]; !ok {
			fmt.Println("not found user")
			return model.UserDefinition{}
		}
		return f.users[id]
	}

	func (f FileDatabase) DeleteUserByID(id string) {
		delete(f.users, id)
	}

	func (f FileDatabase) CreateUser(user model.UserDefinition) {
		if f.users == nil {
			f.users = map[string]model.UserDefinition{"1": {user.Username, user.Password, user.Role}}
			fmt.Println("a")
		} else {
			f.users[string(len(f.users)+1)] = user
			fmt.Println("b")
		}

		//f.users[string(0)] = user
	}

	func (f FileDatabase) UpdateUser(id string, user model.UserDefinition) {
		f.users[id] = user
	}

func (f FileDatabase) AllUsers() map[string]model.UserDefinition {

		return f.users
	}
*/
func (f FileDatabase) Save() error { //това запазва промените във файла на базата данни, не в самата база данни
	//промените се отразяват в този мап, който сме създали
	//а иначе в действителност си става по заявките и в базата данни
	//просто мапът е от юзъри, които поддържат круд операциите и е все едно аналогия на базата
	//и затова няма нужда от тия функции, защото тях ги има за самите юзъри
	//data, err := yaml.Marshal(f.users)
	data, err := json.Marshal(f.users)//подобно - правим си базата във файлов формат
	if err != nil {
		return err
	}
	file, err := os.OpenFile(FileDatabasePath, os.O_CREATE|os.O_TRUNC, 0644)//отваряме файла
	defer file.Close()
	_, err = file.Write(data)//записваме инфото във файла
	return err
}

func (f FileDatabase) Write() map[int]model.UserDefinition {
	f.users = model.Collection

	data, err := yaml.Marshal(f.users)//правим мапа във формат, годен за записване в yaml

	if err != nil {

		log.Fatal(err)
	}

	//fmt.Println(f.users)
	err2 := os.WriteFile(FileDatabasePath, data, 0)//и тук си записваме вече годните данни в yaml формат в yaml файла, който ни е базата

	if err2 != nil {

		log.Fatal(err2)
	}

	fmt.Println("data written")
	return f.users
}
