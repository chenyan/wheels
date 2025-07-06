package dingtalk

import (
	"errors"

	"github.com/chenyan/wheels/httpx/reqs"
)

type Bot struct {
	WebhookURL string
}

type SendResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func NewBot(webhookURL string) *Bot {
	return &Bot{
		WebhookURL: webhookURL,
	}
}

func (b *Bot) Send(msg any) (*SendResult, error) {
	rsp, err := reqs.PostJSON(b.WebhookURL, msg)
	if err != nil {
		return nil, err
	}
	var r SendResult
	if err := rsp.JSON(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (b *Bot) SendText(content string) (*SendResult, error) {
	msg := NewTextMessage()
	msg.SetContent(content)
	return b.Send(msg)
}

func (b *Bot) SendMarkdown(title, text string) (*SendResult, error) {
	msg := NewMarkdownMessage()
	msg.SetTitle(title)
	msg.SetText(text)
	return b.Send(msg)
}

// SendBtnsActionCard 发送带按钮的卡片消息
// title: 卡片标题
// text: 卡片内容
// btns: 按钮列表，必须成对出现, title和actionURL
func (b *Bot) SendBtnsActionCard(title, text string, btns ...string) (*SendResult, error) {
	msg := NewBtnsActionCardMessage()
	msg.SetTitle(title)
	msg.SetText(text)
	if len(btns)%2 != 0 {
		return nil, errors.New("btns must be even")
	}
	for i := 0; i < len(btns); i += 2 {
		msg.AddButton(btns[i], btns[i+1])
	}
	return b.Send(msg)
}