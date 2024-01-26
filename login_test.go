package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamjPriest/HTMXGO/routes"
)

func TestLoginHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/entry", nil) // Assuming login handler is mapped to "/login" path
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	routes.LoginPageHandler(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong Content-Type: got %v want %v", contentType, expectedContentType)
	}

	expectedTemplatePath := "templates/login.html"
	expectedTemplateContent, err := os.ReadFile(expectedTemplatePath)
	if err != nil {
		t.Fatalf("failed to read expected content file: %v", err)
	}

	actualContent := recorder.Body.String()
	if string(expectedTemplateContent) != actualContent {
		t.Errorf("handler returned unexpected body: got %s want %s", actualContent, string(expectedTemplateContent))
	}
}
