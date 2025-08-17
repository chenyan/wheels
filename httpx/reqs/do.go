package reqs

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var (
	DefaultTimeout = 5 * time.Second
	DefaultClient  = &http.Client{Timeout: DefaultTimeout}
)

// Get performs an HTTP GET request to the specified URL and returns a pointer to a Resp struct containing the response and any error encountered.
func Get(url string) *Resp {
	resp, err := DefaultClient.Get(url)
	return &Resp{Response: *resp, Error: err}
}

// PostJSON performs an HTTP POST request to the specified URL with a JSON body and returns a pointer to a Resp struct containing the response and any error encountered.
func PostJSON(url string, body any) (*Resp, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := DefaultClient.Post(url, "application/json", bytes.NewBuffer(bs))
	return &Resp{Response: *resp, Error: err}, nil
}
