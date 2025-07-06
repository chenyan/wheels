package okx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// API 基础 URL
	baseURL = "https://www.okx.com"
	// API 版本
	apiVersion = "v5"
)

// Client OKX API 客户端
type Client struct {
	apiKey     string
	secretKey  string
	passphrase string
	httpClient *http.Client
	baseURL    string
}

// Config 客户端配置
type Config struct {
	APIKey     string
	SecretKey  string
	Passphrase string
	BaseURL    string
}

// NewClient 创建新的 OKX API 客户端
func NewClient(config Config) *Client {
	baseURL := baseURL
	if config.BaseURL != "" {
		baseURL = config.BaseURL
	}

	return &Client{
		apiKey:     config.APIKey,
		secretKey:  config.SecretKey,
		passphrase: config.Passphrase,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// 生成签名
func (c *Client) generateSignature(timestamp, method, requestPath string, body []byte) string {
	message := timestamp + method + requestPath
	if len(body) > 0 {
		message += string(body)
	}

	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 发送请求
func (c *Client) sendRequest(method, path string, body any) ([]byte, error) {
	url := fmt.Sprintf("%s/api/%s%s", c.baseURL, apiVersion, path)
	
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	sign := c.generateSignature(timestamp, method, path, nil)

	req.Header.Set("OK-ACCESS-KEY", c.apiKey)
	req.Header.Set("OK-ACCESS-SIGN", sign)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.passphrase)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(respBody))
	}

	return respBody, nil
}
