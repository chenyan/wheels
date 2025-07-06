package dingtalk

import (
	"encoding/json"
	"testing"
)

func TestBtnsActionCardMessage(t *testing.T) {
	// 创建消息
	msg := NewBtnsActionCardMessage().
		SetTitle("我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身").
		SetText("![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划").
		SetBtnOrientation("0").
		AddButton("内容不错", "https://www.dingtalk.com/").
		AddButton("不感兴趣", "https://www.dingtalk.com/")

	// 序列化为JSON
	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// 反序列化并验证
	var decoded BtnsActionCardMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	// 验证字段
	if decoded.MsgType != "actionCard" {
		t.Errorf("Expected msgtype 'actionCard', got '%s'", decoded.MsgType)
	}

	if decoded.ActionCard.Title != msg.ActionCard.Title {
		t.Errorf("Expected title '%s', got '%s'", msg.ActionCard.Title, decoded.ActionCard.Title)
	}

	if decoded.ActionCard.Text != msg.ActionCard.Text {
		t.Errorf("Expected text '%s', got '%s'", msg.ActionCard.Text, decoded.ActionCard.Text)
	}

	if decoded.ActionCard.BtnOrientation != "0" {
		t.Errorf("Expected btnOrientation '0', got '%s'", decoded.ActionCard.BtnOrientation)
	}

	if len(decoded.ActionCard.Btns) != 2 {
		t.Errorf("Expected 2 buttons, got %d", len(decoded.ActionCard.Btns))
	}

	// 验证按钮
	expectedBtns := []struct {
		title     string
		actionURL string
	}{
		{"内容不错", "https://www.dingtalk.com/"},
		{"不感兴趣", "https://www.dingtalk.com/"},
	}

	for i, expectedBtn := range expectedBtns {
		if decoded.ActionCard.Btns[i].Title != expectedBtn.title {
			t.Errorf("Button %d: expected title '%s', got '%s'", i, expectedBtn.title, decoded.ActionCard.Btns[i].Title)
		}
		if decoded.ActionCard.Btns[i].ActionURL != expectedBtn.actionURL {
			t.Errorf("Button %d: expected actionURL '%s', got '%s'", i, expectedBtn.actionURL, decoded.ActionCard.Btns[i].ActionURL)
		}
	}
}

func TestBtnsActionCardMessage_Chaining(t *testing.T) {
	// 测试方法链式调用
	msg := NewBtnsActionCardMessage()
	
	// 验证每个方法都返回接收者
	if msg.SetTitle("test") != msg {
		t.Error("SetTitle should return receiver")
	}
	if msg.SetText("test") != msg {
		t.Error("SetText should return receiver")
	}
	if msg.SetBtnOrientation("1") != msg {
		t.Error("SetBtnOrientation should return receiver")
	}
	if msg.AddButton("test", "test") != msg {
		t.Error("AddButton should return receiver")
	}
}

func TestTextMessage(t *testing.T) {
	// 创建消息
	msg := NewTextMessage().
		SetContent("我就是我, @014728255240768602 是不一样的烟火").
		AtUsers("014728255240768602")

	// 序列化为JSON
	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// 反序列化并验证
	var decoded TextMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	// 验证字段
	if decoded.MsgType != "text" {
		t.Errorf("Expected msgtype 'text', got '%s'", decoded.MsgType)
	}

	expectedContent := "我就是我, @014728255240768602 是不一样的烟火"
	if decoded.Text.Content != expectedContent {
		t.Errorf("Expected content '%s', got '%s'", expectedContent, decoded.Text.Content)
	}

	if len(decoded.At.AtUserIds) != 1 {
		t.Errorf("Expected 1 atUserId, got %d", len(decoded.At.AtUserIds))
	}

	expectedUserId := "014728255240768602"
	if decoded.At.AtUserIds[0] != expectedUserId {
		t.Errorf("Expected atUserId '%s', got '%s'", expectedUserId, decoded.At.AtUserIds[0])
	}

	if decoded.At.IsAtAll {
		t.Error("Expected IsAtAll to be false")
	}
}

func TestTextMessage_AtAll(t *testing.T) {
	msg := NewTextMessage().
		SetContent("这是一条测试消息").
		AtAll()

	if !msg.At.IsAtAll {
		t.Error("AtAll() failed to set IsAtAll to true")
	}
}

func TestTextMessage_ClearAtUsers(t *testing.T) {
	msg := NewTextMessage().
		AtUsers("user1", "user2").
		AtAll()

	msg.ClearAtUsers()

	if len(msg.At.AtUserIds) != 0 {
		t.Error("ClearAtUsers() failed to clear AtUserIds")
	}

	if msg.At.IsAtAll {
		t.Error("ClearAtUsers() failed to clear IsAtAll")
	}
}

func TestTextMessage_Chaining(t *testing.T) {
	msg := NewTextMessage()

	// 验证每个方法都返回接收者
	if msg.SetContent("test") != msg {
		t.Error("SetContent should return receiver")
	}
	if msg.AtAll() != msg {
		t.Error("AtAll should return receiver")
	}
	if msg.AtUsers("test") != msg {
		t.Error("AtUsers should return receiver")
	}
	if msg.ClearAtUsers() != msg {
		t.Error("ClearAtUsers should return receiver")
	}
}

func TestMarkdownMessage(t *testing.T) {
	// 创建消息
	msg := NewMarkdownMessage().
		SetTitle("杭州天气").
		SetText("#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n").
		AtMobiles("150XXXXXXXX").
		AtUsers("user123")

	// 序列化为JSON
	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	// 反序列化并验证
	var decoded MarkdownMessage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	// 验证字段
	if decoded.MsgType != "markdown" {
		t.Errorf("Expected msgtype 'markdown', got '%s'", decoded.MsgType)
	}

	if decoded.Markdown.Title != "杭州天气" {
		t.Errorf("Expected title '杭州天气', got '%s'", decoded.Markdown.Title)
	}

	expectedText := "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
	if decoded.Markdown.Text != expectedText {
		t.Errorf("Expected text '%s', got '%s'", expectedText, decoded.Markdown.Text)
	}

	if len(decoded.At.AtMobiles) != 1 || decoded.At.AtMobiles[0] != "150XXXXXXXX" {
		t.Error("AtMobiles not set correctly")
	}

	if len(decoded.At.AtUserIds) != 1 || decoded.At.AtUserIds[0] != "user123" {
		t.Error("AtUserIds not set correctly")
	}

	if decoded.At.IsAtAll {
		t.Error("Expected IsAtAll to be false")
	}
}

func TestMarkdownMessage_AtAll(t *testing.T) {
	msg := NewMarkdownMessage().
		SetTitle("测试").
		SetText("测试内容").
		AtAll()

	if !msg.At.IsAtAll {
		t.Error("AtAll() failed to set IsAtAll to true")
	}
}

func TestMarkdownMessage_ClearAt(t *testing.T) {
	msg := NewMarkdownMessage().
		AtMobiles("123", "456").
		AtUsers("user1", "user2").
		AtAll()

	msg.ClearAt()

	if len(msg.At.AtMobiles) != 0 {
		t.Error("ClearAt() failed to clear AtMobiles")
	}

	if len(msg.At.AtUserIds) != 0 {
		t.Error("ClearAt() failed to clear AtUserIds")
	}

	if msg.At.IsAtAll {
		t.Error("ClearAt() failed to clear IsAtAll")
	}
}

func TestMarkdownMessage_Chaining(t *testing.T) {
	msg := NewMarkdownMessage()

	// 验证每个方法都返回接收者
	if msg.SetTitle("test") != msg {
		t.Error("SetTitle should return receiver")
	}
	if msg.SetText("test") != msg {
		t.Error("SetText should return receiver")
	}
	if msg.AtAll() != msg {
		t.Error("AtAll should return receiver")
	}
	if msg.AtMobiles("test") != msg {
		t.Error("AtMobiles should return receiver")
	}
	if msg.AtUsers("test") != msg {
		t.Error("AtUsers should return receiver")
	}
	if msg.ClearAt() != msg {
		t.Error("ClearAt should return receiver")
	}
} 