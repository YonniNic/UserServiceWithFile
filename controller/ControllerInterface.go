package controller

import "net/http"

type Requests interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	ReadEmployees(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	DeleteNewUser(w http.ResponseWriter, r *http.Request)
}
