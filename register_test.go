package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/williamjPriest/HTMXGO/routes"
)

func TestRegisterController(t *testing.T) {
	form := strings.NewReader("username=testuser&password=testpassword")


	req, err := http.NewRequest("POST", "/register", form)
	if err != nil {
		t.Fatal(err)
	}


	recorder := httptest.NewRecorder()

	routes.RegisterHandler(recorder, req)



	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}


type MockDatabase struct{}

func (m *MockDatabase) AddUser(username string, password []byte) error {
	return nil
}
