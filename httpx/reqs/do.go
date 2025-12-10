package reqs

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	DefaultTimeout = 5 * time.Second
	DefaultClient  = &http.Client{}
)

// Get performs an HTTP GET request to the specified URL and returns a pointer to a Resp struct containing the response and any error encountered.
func Get(url string) *Resp {
	return Get2(url, nil)
}

// Get2 performs an HTTP GET request to the specified URL with the given options and returns a pointer to a Resp struct containing the response and any error encountered.
func Get2(url string, opts *Opts) *Resp {
	return Do("GET", url, nil, opts)
}

// PostJSON performs an HTTP POST request to the specified URL with a JSON body and returns a pointer to a Resp struct containing the response and any error encountered.
func PostJSON(url string, body any) *Resp {
	return PostJSON2(url, body, &Opts{})
}

func PostJSON2(url string, body any, opts *Opts) *Resp {
	bs, err := json.Marshal(body)
	if err != nil {
		return &Resp{Error: err}
	}
	if opts == nil {
		opts = &Opts{}
	}
	opts.AddHeader("Content-Type", "application/json")
	return Post2(url, bytes.NewReader(bs), opts)
}

// Post2 performs an HTTP POST request to the specified URL with the given body and options and returns a pointer to a Resp struct containing the response and any error encountered.
func Post2(url string, body io.Reader, opts *Opts) *Resp {
	return Do("POST", url, body, opts)
}

// PostFiles performs an HTTP POST request to the specified URL with a multipart/form-data body and returns a pointer to a Resp struct containing the response and any error encountered.
// The files map is a map of field names and file paths.
// The file paths can start with "@" to indicate a file path.
func PostFiles(url string, files map[string]string) *Resp {
	return PostFiles2(url, files, nil)
}

func PostFiles2(url string, files map[string]string, opts *Opts) *Resp {
	contentType, body, err := createForm(files)
	if err != nil {
		return &Resp{Error: err}
	}
	if opts == nil {
		opts = &Opts{}
	}
	opts.AddHeader("Content-Type", contentType)
	return Post2(url, body, opts)
}

func Do(method string, url string, body io.Reader, opts *Opts) *Resp {
	r, err := BuildRequest(method, url, body, opts)
	if err != nil {
		return &Resp{Error: err}
	}
	resp, err := DefaultClient.Do(r)
	if err != nil {
		return &Resp{Error: err}
	}
	return &Resp{Response: *resp, Error: nil}
}
