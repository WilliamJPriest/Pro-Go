package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/williamjPriest/HTMXGO/routes"
)

func TestRegisterHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	routes.RegisterHandler(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong Content-Type: got %v want %v", contentType, expectedContentType)
	}

	expectedTemplateContent := "Expected Content in login.html"
	if body := recorder.Body.String(); !strings.Contains(body, expectedTemplateContent) {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedTemplateContent)
	}
}