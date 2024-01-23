package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williamjPriest/HTMXGO/routes"
)

func TestRegisterPageHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/registerForm", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	routes.RegisterPageHandler(rr, req)


	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

}