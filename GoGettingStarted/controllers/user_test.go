package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_userController_PostUser(t *testing.T) {

	payload := strings.NewReader(`{
    "FirstName" : "Shrikant",
    "LastName" : "Urkande"
}`)
	req, err := http.NewRequest("POST", "/users", payload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	uc := newUserController()
	// handler := http.HandlerFunc(uc)
	uc.PostUser(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ID":1,"FirstName":"Shrikant","LastName":"Urkande"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_userController_GetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	uc := newUserController()
	// handler := http.HandlerFunc(uc)
	uc.GetAllUsers(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":1,"FirstName":"Shrikant","LastName":"Urkande"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
