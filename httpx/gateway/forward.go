package gateway

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

var (
	Logger *slog.Logger
)

// Forward forwards an HTTP request to a target endpoint and writes the response to the provided ResponseWriter.
// It logs the request method, raw path, response size, and any error that occurs during the process.
//
// Parameters:
//   - r: The original HTTP request to be forwarded.
//   - targetEndpoint: The target endpoint to which the request should be forwarded.
//   - path: The path to be appended to the target endpoint.
//   - w: The ResponseWriter to which the response should be written.
//
// Returns:
//   - error: An error if any occurs during the forwarding process.
func Forward(r *http.Request, targetEndpoint, path string, w *http.ResponseWriter) error {
	var (
		rawPath string
		rspSize int64
		errx    error
	)
	defer func() {
		Logger.DebugContext(
			r.Context(), "forword", "method", r.Method, "rawpath", rawPath, "rspsz", rspSize, "err", errx)
	}()
	url := fmt.Sprintf("%s%s", targetEndpoint, path)
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		errx = err
		return err
	}
	for k, v := range r.Header {
		req.Header[k] = v
	}
	rsp, err := (&http.Client{}).Do(req)
	if err != nil {
		errx = err
		return err
	}
	defer rsp.Body.Close()

	// update writer header
	UpdateHeader(w, rsp.Header)

	rspSize, err = io.Copy(*w, rsp.Body)
	return err
}

func ForwardStream(r *http.Request, targetEndpoint, path string, w *http.ResponseWriter) error {
	var (
		rawPath string
		errx    error
	)
	defer func() {
		Logger.DebugContext(
			r.Context(), "forwordstream", "method", r.Method, "rawpath", rawPath, "err", errx)
	}()
	url := fmt.Sprintf("%s%s", targetEndpoint, path)
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		errx = err
		return err
	}
	for k, v := range r.Header {
		req.Header[k] = v
	}
	rsp, err := (&http.Client{}).Do(req)
	if err != nil {
		errx = err
		return err
	}
	defer rsp.Body.Close()

	// update writer header
	UpdateHeader(w, rsp.Header)

	reader := bufio.NewReader(rsp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		(*w).Write(line)
		(*w).(http.Flusher).Flush()
	}
	return err
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
