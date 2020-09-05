package main_test

import (
	"net/http"
	"testing"
)

func TestGetOriginalURL(t *testing.T) {
	// make a dummy reques
	response, err := http.Get("http://localhost:8080/8n90N2")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Error("Encountered an error:", err)
	}
}

func TestGetOriginalURL2(t *testing.T) {
	// make a dummy reques
	response, err := http.Get("http://localhost:8080/MEhRlZ")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Error("Encountered an error:", err)
	}
}
