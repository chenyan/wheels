package reqs

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestPostJSON(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		bs, err := io.ReadAll(req.Body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
		}
		rw.Write(bs)
	}))
	defer server.Close()

	// Make a POST request to the local server
	resp := PostJSON(server.URL, map[string]any{"a": 1, "b": "B"})

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

func TestPostFiles(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(1024 * 1024)
		form := req.MultipartForm
		if form == nil {
			t.Errorf("Expected multipart form to be not nil")
		}
		if len(form.File) == 0 {
			t.Errorf("Expected file to be not nil")
		}
		if form.File["file"] == nil {
			t.Errorf("Expected file to be not nil")
		}
		log.Println(form)
		file, err := form.File["file"][0].Open()
		if err != nil {
			t.Errorf("Error opening file: %v", err)
		}
		defer file.Close()
		bs, err := io.ReadAll(file)
		if err != nil {
			t.Errorf("Error reading file: %v", err)
		}
		if string(bs) != "Hello, world!" {
			t.Errorf("Expected file content to be 'Hello, world!', but got '%s'", string(bs))
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"size": ` + strconv.Itoa(len(bs)) + `}`))
	}))
	defer server.Close()

	// Make a POST request to the local server
	resp := PostFiles(server.URL, map[string]string{"a": "1", "file": "@testdata/test.txt"})
	mapData, err := resp.JSONMap()
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}
	if mapData["size"] == 0 {
		t.Errorf("Expected size to be 12, but got %d", mapData["size"])
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}
