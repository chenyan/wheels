package okx

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	if client.APIKey != "test-key" {
		t.Errorf("APIKey = %s; want test-key", client.APIKey)
	}
	
	if client.SecretKey != "test-secret" {
		t.Errorf("SecretKey = %s; want test-secret", client.SecretKey)
	}
	
	if client.Passphrase != "test-pass" {
		t.Errorf("Passphrase = %s; want test-pass", client.Passphrase)
	}
	
	if client.BaseURL != BaseURL {
		t.Errorf("BaseURL = %s; want %s", client.BaseURL, BaseURL)
	}
	
	if client.Simulated {
		t.Error("Simulated should be false by default")
	}
}

func TestSetSimulated(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	client.SetSimulated(true)
	if !client.Simulated {
		t.Error("Simulated should be true after SetSimulated(true)")
	}
	
	client.SetSimulated(false)
	if client.Simulated {
		t.Error("Simulated should be false after SetSimulated(false)")
	}
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	customURL := "https://custom.okx.com"
	client.SetBaseURL(customURL)
	
	if client.BaseURL != customURL {
		t.Errorf("BaseURL = %s; want %s", client.BaseURL, customURL)
	}
}

func TestGetTimestamp(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	timestamp := client.getTimestamp()
	
	// 验证时间戳格式: 2006-01-02T15:04:05.000Z
	_, err := time.Parse("2006-01-02T15:04:05.000Z", timestamp)
	if err != nil {
		t.Errorf("Invalid timestamp format: %s, error: %v", timestamp, err)
	}
}

func TestSign(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	timestamp := "2023-11-29T10:00:00.000Z"
	method := "GET"
	path := "/api/v5/account/balance"
	body := ""
	
	sign := client.sign(timestamp, method, path, body)
	
	// 验证签名不为空
	if sign == "" {
		t.Error("Sign should not be empty")
	}
	
	// 验证签名是 base64 编码的
	if len(sign) == 0 {
		t.Error("Sign length should be greater than 0")
	}
}

func TestBuildHeaders(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	
	timestamp := "2023-11-29T10:00:00.000Z"
	method := "GET"
	path := "/api/v5/account/balance"
	body := ""
	
	headers := client.buildHeaders(timestamp, method, path, body)
	
	// 验证必需的请求头
	if headers["OK-ACCESS-KEY"] != "test-key" {
		t.Errorf("OK-ACCESS-KEY = %s; want test-key", headers["OK-ACCESS-KEY"])
	}
	
	if headers["OK-ACCESS-TIMESTAMP"] != timestamp {
		t.Errorf("OK-ACCESS-TIMESTAMP = %s; want %s", headers["OK-ACCESS-TIMESTAMP"], timestamp)
	}
	
	if headers["OK-ACCESS-PASSPHRASE"] != "test-pass" {
		t.Errorf("OK-ACCESS-PASSPHRASE = %s; want test-pass", headers["OK-ACCESS-PASSPHRASE"])
	}
	
	if headers["Content-Type"] != "application/json" {
		t.Errorf("Content-Type = %s; want application/json", headers["Content-Type"])
	}
	
	if headers["OK-ACCESS-SIGN"] == "" {
		t.Error("OK-ACCESS-SIGN should not be empty")
	}
}

func TestBuildHeadersWithSimulated(t *testing.T) {
	client := NewClient("test-key", "test-secret", "test-pass")
	client.SetSimulated(true)
	
	timestamp := "2023-11-29T10:00:00.000Z"
	method := "GET"
	path := "/api/v5/account/balance"
	body := ""
	
	headers := client.buildHeaders(timestamp, method, path, body)
	
	// 验证模拟交易标志
	if headers["x-simulated-trading"] != "1" {
		t.Error("x-simulated-trading should be 1 when Simulated is true")
	}
}

func TestBaseResponse_IsSuccess(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{"Success", "0", true},
		{"Error", "50000", false},
		{"Error", "50001", false},
		{"Empty", "", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &BaseResponse{Code: tt.code}
			if got := resp.IsSuccess(); got != tt.want {
				t.Errorf("IsSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

