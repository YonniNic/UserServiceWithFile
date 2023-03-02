package repository

import (
	"encoding/json"
	"os"
)

type DataBase interface {
	/*GetUserByID(id string) model.UserDefinition
	DeleteUserByID(id string)
	CreateUser(user model.UserDefinition)
	UpdateUser(id string, user model.UserDefinition)
	AllUsers() map[string]model.UserDefinition
	*/
}

func NewFileDataBaseDefinition() (FileDatabase, error) { //това създава нов файл да се запише инфото за базата данни, не нова база данни
	file, err := os.OpenFile(FileDatabasePath, os.O_CREATE, 0644)
	defer file.Close()

	var fileData []byte
	file.Read(fileData)

	data := FileDatabase{}
	json.Unmarshal(fileData, data)
	return data, err
}
