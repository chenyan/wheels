package reqs

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Opts struct {
	Timeout        time.Duration
	Headers        map[string]string
	Cookies        map[string]string
	Query          url.Values
	Stream         bool
	AllowRedirects bool

	cookies []*http.Cookie
}

// createForm creates a multipart/form-data body from a map of strings.
// The keys are the field names, and the values are the field values.
// If the value starts with "@", it is treated as a file path.
// The file is read and added to the multipart/form-data body.
// The function returns the content type and the body reader.
// The content type is "multipart/form-data; boundary=...".
func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	for key, value := range form {
		if strings.HasPrefix(value, "@") {
			file, err := os.Open(value[1:])
			if err != nil {
				return "", nil, fmt.Errorf("failed to open file: %s %w", value[1:], err)
			}
			defer file.Close()
			filename := filepath.Base(value[1:])
			fw, err := mw.CreateFormFile(key, filename)
			if err != nil {
				return "", nil, fmt.Errorf("failed to create form file: %s %w", key, err)
			}
			_, err = io.Copy(fw, file)
			if err != nil {
				return "", nil, fmt.Errorf("failed to copy file: %s %w", key, err)
			}
		} else {
			mw.WriteField(key, value)
		}
	}
	err := mw.Close()
	if err != nil {
		return "", nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}
	return mw.FormDataContentType(), body, nil
}

func (o *Opts) AddCookie(cookie *http.Cookie) *Opts {
	o.cookies = append(o.cookies, cookie)
	return o
}

// BuildRequest 创建一个http请求, 合并opts中的参数到请求中, 返回请求对象和错误
func BuildRequest(method string, urlstr string, body io.Reader, opts *Opts) (*http.Request, error) {
	if opts == nil {
		return http.NewRequest(method, urlstr, body)
	}

	if opts.Cookies != nil {
		for name, value := range opts.Cookies {
			cookie := &http.Cookie{Name: name, Value: value}
			opts.AddCookie(cookie)
		}
	}

	if len(opts.Query) > 0 {
		u, err := url.Parse(urlstr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse url: %s %w", urlstr, err)
		}
		if len(u.Query()) > 0 {
			for key, values := range u.Query() {
				for _, value := range values {
					opts.Query.Add(key, value)
				}
			}
		}
		u.RawQuery = opts.Query.Encode()
		urlstr = u.String()
	}

	r, err := http.NewRequest(method, urlstr, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s %w", urlstr, err)
	}

	for key, value := range opts.Headers {
		r.Header.Add(key, value)
	}
	return r, nil
}

func Q(kv ...string) url.Values {
	if len(kv)%2 != 0 {
		panic("odd number of arguments")
	}
	values := url.Values{}
	for i := 0; i < len(kv); i += 2 {
		values.Add(kv[i], kv[i+1])
	}
	return values
}
