package main

import (
	"UserService/model"
	"UserService/repository"
	"UserService/routers"
	"UserService/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	database, err := repository.NewFileDataBaseDefinition()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pannel := service.NewUserPanel(model.UserDefinition{"", "", "Admin"}, database)
	user := model.UserDefinition{Username: "hello", Password: "you", Role: "there"}
	user1 := model.UserDefinition{Username: "bye", Password: "you", Role: "there"}
	user2 := model.UserDefinition{Username: "bye", Password: "you", Role: "there"}
	user3 := model.UserDefinition{Username: "goodmorning", Password: "you", Role: "there"}
	user4 := model.UserDefinition{Username: "success", Password: "you", Role: "there"}
	user5 := model.UserDefinition{Username: "great", Password: "do", Role: "job"}
	//user10 := model.UserDefinition{Username: "user", Password: "user", Role: "user"}

	pannel.Create(user)
	pannel.Create(user1)
	pannel.Create(user2)
	pannel.Create(user3)
	pannel.Create(user4)
	pannel.Create(user5)

	fmt.Println(pannel.Logging("bye", "you"))

	fmt.Println(pannel.ViewDataForUser(2))
	pannel.DeleteUser(2)
	//pannel.Delete(6)
	database.Write()
	fmt.Println(pannel.GetUserById(5))

	//testing.TestAGetUserById()

	//fmt.Println(pannel.GetUserById("\u0005"))
	//fmt.Println(pannel.AllUsers())

	router := mux.NewRouter().StrictSlash(true)
	routers.BasicCrudRoutes(router)

	//database.Save()
	http.ListenAndServe(":8080", router)

}
