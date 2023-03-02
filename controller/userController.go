package controller

import (
	"UserService/model"
	"UserService/repository"
	"UserService/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"text/template"
	"time"
	//"gopkg.in/yaml.v3"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("/Users/I556357/Downloads/Task-main 2/view/index.tmpl"))

type Str struct {
	user model.UserDefinition
	f    repository.FileDatabase
}

type Pannel struct {
	pann service.UserPanel
}

func (u Str) Index(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Index", nil)
}

func (u Str) Hello(w http.ResponseWriter, r *http.Request) {
	newUser := model.UserDefinition{r.PostFormValue("username"),
		r.PostFormValue("password"),
		r.PostFormValue("role")}

	u.user.Create(newUser)
	u.f.Write()
	json.NewEncoder(w).Encode(newUser)

}

func (u Str) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	user := &model.UserDefinition{}
	json.NewDecoder(r.Body).Decode(&user)
	u.user.Create(*user)
	u.f.Write()
	json.NewEncoder(w).Encode(user)
	tmpl.ExecuteTemplate(w, "Success", nil)
}

func (u Str) Loginbrowser(w http.ResponseWriter, r *http.Request) {
	user := model.UserDefinition{Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password")}

	//json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
	var expectedPassword string

	for i := 0; i <= len(model.Collection); i++ {
		if user.Username == model.Collection[i].Username {
			expectedPassword = model.Collection[i].Password
		}

	}

	if expectedPassword != user.Password {
		json.NewEncoder(w).Encode("Wrong password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: user.Username,
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	c := &http.Cookie{
		Name:    "1",
		Value:   sessionToken,
		Expires: expiresAt,
	}
	r.AddCookie(c)

	c, err := r.Cookie("1")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken = c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}

func (p Pannel) Login(w http.ResponseWriter, r *http.Request) {
	user := &model.UserDefinition{}
	json.NewDecoder(r.Body).Decode(&user)
	var expectedPassword string
	var role string

	id, _ := p.pann.Logging(user.Username, user.Password)

	expectedPassword = model.Collection[id].Password
	role = model.Collection[id].Role
	user.Role = role
	json.NewEncoder(w).Encode(user)

	if expectedPassword != user.Password || user.Password == "" {
		json.NewEncoder(w).Encode("Invalid input")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(180 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: user.Username,
		expiry:   expiresAt,
	}

	json.NewEncoder(w).Encode(len(sessions))

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	c := &http.Cookie{
		Name:    "1",
		Value:   sessionToken,
		Expires: expiresAt,
	}
	r.AddCookie(c)

	json.NewEncoder(w).Encode("yes")
	c, err := r.Cookie("1")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken = c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}

/*
	func Welcome(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		// We then get the session from our session map
		userSession, exists := sessions[sessionToken]
		if !exists {
			// If the session token is not present in session map, return an unauthorized error
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the session is present, but has expired, we can delete the session, and return
		// an unauthorized status
		if userSession.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// If the session is valid, return the welcome message to the user
		w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
	}
*/
func (u Str) ReadUsers(w http.ResponseWriter, r *http.Request) {
	var users map[int]model.UserDefinition
	users = u.user.Users()
	json.NewEncoder(w).Encode(users)

}
func (u Str) GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])
	var user model.UserDefinition
	user = u.user.GetUserById(id)
	//json.NewEncoder(w).Encode(id)
	json.NewEncoder(w).Encode(user)
}

func (u Str) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])
	user := u.user.GetUserById(id)
	newUser := &model.UserDefinition{}
	json.NewDecoder(r.Body).Decode(&newUser)
	user = u.user.Update(id, *newUser)
	json.NewEncoder(w).Encode(user)
}

func (u Str) DeleteNewUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])
	user := u.user.Delete(id)
	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode("The User is Deleted Successfully!")
}
