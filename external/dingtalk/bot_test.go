package dingtalk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupTestServer(t *testing.T, expectedMsg interface{}, result *SendResult) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// 验证 Content-Type
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", ct)
		}

		// 解码请求体
		var receivedMsg map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&receivedMsg); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		// 将预期消息转换为相同的格式进行比较
		expectedJSON, _ := json.Marshal(expectedMsg)
		var expectedMap map[string]interface{}
		json.Unmarshal(expectedJSON, &expectedMap)

		// 使用 reflect.DeepEqual 进行深度比较
		if !reflect.DeepEqual(expectedMap, receivedMsg) {
			expectedPretty, _ := json.MarshalIndent(expectedMap, "", "  ")
			receivedPretty, _ := json.MarshalIndent(receivedMsg, "", "  ")
			t.Errorf("Messages don't match.\nExpected:\n%s\nGot:\n%s", expectedPretty, receivedPretty)
		}

		// 返回结果
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}))
}

func TestBot_Send(t *testing.T) {
	// 准备测试数据
	msg := NewTextMessage().SetContent("测试消息")
	result := &SendResult{
		ErrCode: 0,
		ErrMsg:  "ok",
	}

	// 创建测试服务器
	server := setupTestServer(t, msg, result)
	defer server.Close()

	// 创建 Bot 实例
	bot := NewBot(server.URL)

	// 发送消息
	resp, err := bot.Send(msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// 验证响应
	if resp.ErrCode != result.ErrCode {
		t.Errorf("Expected error code %d, got %d", result.ErrCode, resp.ErrCode)
	}
	if resp.ErrMsg != result.ErrMsg {
		t.Errorf("Expected error message %s, got %s", result.ErrMsg, resp.ErrMsg)
	}
}

func TestBot_SendText(t *testing.T) {
	content := "测试消息"
	msg := NewTextMessage().SetContent(content)
	result := &SendResult{ErrCode: 0, ErrMsg: "ok"}

	server := setupTestServer(t, msg, result)
	defer server.Close()

	bot := NewBot(server.URL)
	resp, err := bot.SendText(content)
	if err != nil {
		t.Fatalf("Failed to send text message: %v", err)
	}

	if resp.ErrCode != 0 {
		t.Errorf("Expected error code 0, got %d", resp.ErrCode)
	}
}

func TestBot_SendMarkdown(t *testing.T) {
	title := "测试标题"
	text := "测试内容"
	msg := NewMarkdownMessage().SetTitle(title).SetText(text)
	result := &SendResult{ErrCode: 0, ErrMsg: "ok"}

	server := setupTestServer(t, msg, result)
	defer server.Close()

	bot := NewBot(server.URL)
	resp, err := bot.SendMarkdown(title, text)
	if err != nil {
		t.Fatalf("Failed to send markdown message: %v", err)
	}

	if resp.ErrCode != 0 {
		t.Errorf("Expected error code 0, got %d", resp.ErrCode)
	}
}

func TestBot_SendBtnsActionCard(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		text    string
		btns    []string
		wantErr bool
	}{
		{
			name:    "正常按钮",
			title:   "测试标题",
			text:    "测试内容",
			btns:    []string{"按钮1", "URL1", "按钮2", "URL2"},
			wantErr: false,
		},
		{
			name:    "奇数个按钮参数",
			title:   "测试标题",
			text:    "测试内容",
			btns:    []string{"按钮1", "URL1", "按钮2"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewBtnsActionCardMessage().SetTitle(tt.title).SetText(tt.text)
			if !tt.wantErr {
				for i := 0; i < len(tt.btns); i += 2 {
					msg.AddButton(tt.btns[i], tt.btns[i+1])
				}
			}
			result := &SendResult{ErrCode: 0, ErrMsg: "ok"}

			server := setupTestServer(t, msg, result)
			defer server.Close()

			bot := NewBot(server.URL)
			resp, err := bot.SendBtnsActionCard(tt.title, tt.text, tt.btns...)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Failed to send action card message: %v", err)
			}

			if resp.ErrCode != 0 {
				t.Errorf("Expected error code 0, got %d", resp.ErrCode)
			}
		})
	}
}

func TestNewBot(t *testing.T) {
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=xxx"
	bot := NewBot(webhookURL)

	if bot == nil {
		t.Error("NewBot() returned nil")
	}

	if bot.WebhookURL != webhookURL {
		t.Errorf("Expected webhook URL %s, got %s", webhookURL, bot.WebhookURL)
	}
} 