package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamjPriest/HTMXGO/routes"
)

func TestRegisterPageHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/registerForm", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	routes.RegisterPageHandler(recorder, req)


	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	
	expectedTemplatePath := "templates/register.html"
	expectedTemplateContent, err := os.ReadFile(expectedTemplatePath)
	if err != nil {
		t.Fatalf("failed to read expected content file: %v", err)
	}

	actualContent := recorder.Body.String()
	if string(expectedTemplateContent) != actualContent {
		t.Errorf("handler returned unexpected body: got %s want %s", actualContent, string(expectedTemplateContent))
	}

}