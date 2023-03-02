package routers

import (
	"UserService/controller"
	"github.com/gorilla/mux"
)

func BasicCrudRoutes(router *mux.Router) {
	/*http.HandleFunc("/user/aaa", controller.Str{}.CreateNewUser)
	http.HandleFunc("/api/user", controller.Str{}.ReadEmployees)
	http.HandleFunc("/user/{update/id}", controller.Str{}.UpdateUser)
	http.HandleFunc("/user/{id}", controller.Str{}.GetUser)
	http.HandleFunc("/user/delete/{id}", controller.Str{}.DeleteNewUser)
	*/
	router.HandleFunc("/api/users/get/{id}", controller.Str{}.GetUser)
	router.HandleFunc("/api/users", controller.Str{}.ReadUsers)
	router.HandleFunc("/api/users/create", controller.Str{}.CreateNewUser)
	router.HandleFunc("/api/login", controller.Pannel{}.Login)
	router.HandleFunc("/api/users/update/{id}", controller.Str{}.UpdateUser)
	router.HandleFunc("/api/users/delete/{id}", controller.Str{}.DeleteNewUser)

	router.HandleFunc("/api/", controller.Str{}.Index)
	router.HandleFunc("/api/hello", controller.Str{}.Hello)
	router.HandleFunc("/api/signin", controller.Str{}.Loginbrowser)
	//router.HandleFunc("/api/welcome", controller.Welcome)
}

