package gateway

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

var (
	Logger *slog.Logger
)

// Forward forwards the request to targetEndpoint + path
func Forward(r *http.Request, targetEndpoint, path string) (http.Header, []byte, error) {
	var (
		rawPath string
		reqSize int
		rspSize int
		errx    error
	)
	defer func() {
		Logger.DebugContext(
			r.Context(), "forword", "method", r.Method, "rawpath", rawPath, "reqsz", reqSize, "rspsz", rspSize, "err", errx)
	}()
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		errx = err
		return nil, nil, err
	}
	reqSize = len(rBody)
	url := fmt.Sprintf("%s%s", targetEndpoint, path)
	req, err := http.NewRequest(r.Method, url, bytes.NewReader(rBody))
	if err != nil {
		errx = err
		return nil, nil, err
	}
	for k, v := range r.Header {
		req.Header[k] = v
	}
	rsp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer rsp.Body.Close()
	rspBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		errx = err
		return nil, nil, err
	}
	rspSize = len(rspBody)
	return rsp.Header, rspBody, nil
}

// UpdateHeader updates the header of the response writer
func UpdateHeader(w *http.ResponseWriter, header http.Header) {
	h := (*w).Header()
	for k, v := range header {
		if len(v) == 0 {
			continue
		}
		h.Set(k, v[0])
	}
}
