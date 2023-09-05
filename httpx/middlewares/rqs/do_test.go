package rqs

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))
	defer server.Close()

	// Make a GET request to the local server
	resp := Get(server.URL)

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check that the response body matches the expected value
	expectedBody := "Hello, world!"
	actualBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(actualBody) != expectedBody {
		t.Errorf("Expected response body '%s', but got '%s'", expectedBody, string(actualBody))
	}
}

func TestGetJSON(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"a": 1, "b": "B"}`))
	}))
	defer server.Close()

	// Make a GET request to the local server
	resp := Get(server.URL)

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	var data struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	err := resp.JSON(&data)
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}
	defer resp.Body.Close()

	// Check that the response body matches the expected value
	if data.A != 1 || data.B != "B" {
		t.Errorf("Expected JSON data to be {1, \"B\"}, but got %v", data)
	}
}
