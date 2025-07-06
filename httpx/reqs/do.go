package reqs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Get performs an HTTP GET request to the specified URL and returns a pointer to a Resp struct containing the response and any error encountered.
func Get(url string) *Resp {
	resp, err := http.Get(url)
	return &Resp{Response: *resp, Error: err}
}

// PostJSON performs an HTTP POST request to the specified URL with a JSON body and returns a pointer to a Resp struct containing the response and any error encountered.
func PostJSON(url string, body any) (*Resp, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	return &Resp{Response: *resp, Error: nil}, nil
}

