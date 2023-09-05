package rqs

import "net/http"

// Get performs an HTTP GET request to the specified URL and returns a pointer to a Resp struct containing the response and any error encountered.
func Get(url string) *Resp {
	resp, err := http.Get(url)
	return &Resp{Response: *resp, Error: err}
}
