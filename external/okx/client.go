package okx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/chenyan/wheels/httpx/reqs"
)

const (
	// BaseURL OKX API 基础 URL
	BaseURL = "https://www.okx.com"
	// BaseURLAWS OKX AWS 服务器 URL
	BaseURLAWS = "https://aws.okx.com"
)

// Client OKX API 客户端
type Client struct {
	APIKey     string
	SecretKey  string
	Passphrase string
	BaseURL    string
	Simulated  bool // 是否使用模拟交易
}

// NewClient 创建一个新的 OKX API 客户端
func NewClient(apiKey, secretKey, passphrase string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		BaseURL:    BaseURL,
		Simulated:  false,
	}
}

// SetSimulated 设置是否使用模拟交易
func (c *Client) SetSimulated(simulated bool) *Client {
	c.Simulated = simulated
	return c
}

// SetBaseURL 设置基础 URL
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.BaseURL = baseURL
	return c
}

// sign 生成请求签名
func (c *Client) sign(timestamp, method, requestPath, body string) string {
	message := timestamp + method + requestPath + body
	h := hmac.New(sha256.New, []byte(c.SecretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// getTimestamp 获取 ISO 8601 格式的时间戳
func (c *Client) getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}

// buildHeaders 构建请求头
func (c *Client) buildHeaders(timestamp, method, requestPath, body string) map[string]string {
	sign := c.sign(timestamp, method, requestPath, body)
	headers := map[string]string{
		"OK-ACCESS-KEY":        c.APIKey,
		"OK-ACCESS-SIGN":       sign,
		"OK-ACCESS-TIMESTAMP":  timestamp,
		"OK-ACCESS-PASSPHRASE": c.Passphrase,
		"Content-Type":         "application/json",
	}

	// 如果是模拟交易，添加相应的标志
	if c.Simulated {
		headers["x-simulated-trading"] = "1"
	}

	return headers
}

// Request 发送请求
func (c *Client) Request(method, path string, body interface{}) (*reqs.Resp, error) {
	var bodyReader io.Reader
	var bodyString string

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body error: %w", err)
		}
		bodyString = string(jsonBody)
		bodyReader = bytes.NewReader(jsonBody)
	}

	timestamp := c.getTimestamp()
	headers := c.buildHeaders(timestamp, method, path, bodyString)

	url := c.BaseURL + path

	opts := &reqs.Opts{
		Headers: headers,
		Timeout: 10 * time.Second,
	}

	resp := reqs.Do(method, url, bodyReader, opts)
	return resp, resp.Error
}

// Get 发送 GET 请求
func (c *Client) Get(path string) (*reqs.Resp, error) {
	return c.Request("GET", path, nil)
}

// Post 发送 POST 请求
func (c *Client) Post(path string, body interface{}) (*reqs.Resp, error) {
	return c.Request("POST", path, body)
}

// BaseResponse OKX API 基础响应结构
type BaseResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data,omitempty"`
}

// IsSuccess 判断响应是否成功
func (r *BaseResponse) IsSuccess() bool {
	return r.Code == "0"
}

// ParseResponse 解析响应
func ParseResponse(resp *reqs.Resp, data interface{}) error {
	if resp.Error != nil {
		return resp.Error
	}
	defer resp.Close()

	var baseResp BaseResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body error: %w", err)
	}

	if err := json.Unmarshal(bodyBytes, &baseResp); err != nil {
		return fmt.Errorf("unmarshal response error: %w", err)
	}

	if !baseResp.IsSuccess() {
		return fmt.Errorf("api error: code=%s, msg=%s", baseResp.Code, baseResp.Msg)
	}

	if data != nil && len(baseResp.Data) > 0 {
		if err := json.Unmarshal(baseResp.Data, data); err != nil {
			return fmt.Errorf("unmarshal data error: %w", err)
		}
	}

	return nil
}
