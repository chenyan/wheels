package reqs

import (
	"encoding/json"
	"io"
	"net/http"
)

// Resp is the response of a request
type Resp struct {
	http.Response
	Error error
}

func (r *Resp) JSON(v any) error {
	if r.Error != nil {
		return r.Error
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func (r *Resp) JSONMap() (map[string]any, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	var v map[string]any
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}

func (r *Resp) String() (string, error) {
	if r.Error != nil {
		return "", r.Error
	}
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bs), err
}
