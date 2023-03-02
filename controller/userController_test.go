package controller
import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var u Str

func TestCreateNewUser(t *testing.T) {
	var jsonStr = [][]byte{[]byte(`{"USERNAME":"test","PASSWORD":"test","ROLE":"test"}`),
		[]byte(`{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}`),
		[]byte(`{"USERNAME":"test2","PASSWORD":"test2","ROLE":"test2"}`),
		[]byte(`{"USERNAME":"test3","PASSWORD":"test3","ROLE":"test3"}`),
	}

	req, err := http.NewRequest("POST", "/api/users/create", bytes.NewBuffer(jsonStr[0]))
	req1, err := http.NewRequest("POST", "/api/users/create", bytes.NewBuffer(jsonStr[1]))
	req2, err := http.NewRequest("POST", "/api/users/create", bytes.NewBuffer(jsonStr[2]))
	req3, err := http.NewRequest("POST", "/api/users/create", bytes.NewBuffer(jsonStr[3]))

	//req, err = http.NewRequest("POST", "/api/users/create", bytes.NewBuffer(jsonStr1))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	u.CreateNewUser(rr, req)
	u.CreateNewUser(rr, req1)
	u.CreateNewUser(rr, req2)
	u.CreateNewUser(rr, req3)

	//fmt.Println(u.user.Users())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"USERNAME":"test","PASSWORD":"test","ROLE":"test"}` + "\n" +
		`{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}` + "\n" +
		`{"USERNAME":"test2","PASSWORD":"test2","ROLE":"test2"}` + "\n" +
		`{"USERNAME":"test3","PASSWORD":"test3","ROLE":"test3"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAll(t *testing.T) {
	TestCreateNewUser(t)
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	//fmt.Println(req)
	handler := http.HandlerFunc(u.ReadUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"1":{"USERNAME":"test","PASSWORD":"test","ROLE":"test"}` + "," +
		`"2":{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}` + "," +
		`"3":{"USERNAME":"test2","PASSWORD":"test2","ROLE":"test2"}` + "," +
		`"4":{"USERNAME":"test3","PASSWORD":"test3","ROLE":"test3"}}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdate(t *testing.T) {
	TestCreateNewUser(t)
	var jsonStr = []byte(`{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}`)

	req, err := http.NewRequest("PUT", "/api/users/update/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(u.UpdateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDelete(t *testing.T) {
	TestCreateNewUser(t)
	req, err := http.NewRequest("DELETE", "/api/users/delete/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(u.DeleteNewUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"USERNAME":"test","PASSWORD":"test","ROLE":"test"}` + "\n" + `"The User is Deleted Successfully!"` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	TestCreateNewUser(t)
	req, err := http.NewRequest("GET", "/api/users/get/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "2",
	}

	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()

	u.GetUser(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"USERNAME":"test1","PASSWORD":"test1","ROLE":"test1"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
