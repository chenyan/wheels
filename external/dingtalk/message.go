package dingtalk

// BtnsActionCardMessage 表示一个带按钮的卡片消息
type BtnsActionCardMessage struct {
	MsgType    string           `json:"msgtype"`
	ActionCard BtnsActionCard   `json:"actionCard"`
}

// BtnsActionCard 表示卡片内容
type BtnsActionCard struct {
	Title          string         `json:"title"`
	Text           string         `json:"text"` // 支持markdown
	BtnOrientation string         `json:"btnOrientation"`
	Btns           []ActionButton `json:"btns"`
}

// ActionButton 表示卡片中的按钮
type ActionButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// NewBtnsActionCardMessage 创建一个新的按钮式卡片消息
func NewBtnsActionCardMessage() *BtnsActionCardMessage {
	return &BtnsActionCardMessage{
		MsgType: "actionCard",
		ActionCard: BtnsActionCard{
			BtnOrientation: "0", // 默认按钮水平排列
		},
	}
}

// SetTitle 设置卡片标题
func (m *BtnsActionCardMessage) SetTitle(title string) *BtnsActionCardMessage {
	m.ActionCard.Title = title
	return m
}

// SetText 设置卡片内容
func (m *BtnsActionCardMessage) SetText(text string) *BtnsActionCardMessage {
	m.ActionCard.Text = text
	return m
}

// SetBtnOrientation 设置按钮排列方向
// orientation: "0"-按钮水平排列，"1"-按钮垂直排列
func (m *BtnsActionCardMessage) SetBtnOrientation(orientation string) *BtnsActionCardMessage {
	m.ActionCard.BtnOrientation = orientation
	return m
}

// AddButton 添加按钮
func (m *BtnsActionCardMessage) AddButton(title, actionURL string) *BtnsActionCardMessage {
	m.ActionCard.Btns = append(m.ActionCard.Btns, ActionButton{
		Title:     title,
		ActionURL: actionURL,
	})
	return m
}

// TextMessage 表示文本消息
type TextMessage struct {
	MsgType string    `json:"msgtype"`
	Text    TextContent `json:"text"`
	At      AtContent  `json:"at"`
}

// TextContent 表示文本消息内容
type TextContent struct {
	Content string `json:"content"`
}

// AtContent 表示@相关设置
type AtContent struct {
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

// NewTextMessage 创建一个新的文本消息
func NewTextMessage() *TextMessage {
	return &TextMessage{
		MsgType: "text",
		Text:    TextContent{},
		At:      AtContent{AtUserIds: make([]string, 0)},
	}
}

// SetContent 设置消息内容
func (m *TextMessage) SetContent(content string) *TextMessage {
	m.Text.Content = content
	return m
}

// AtAll @所有人
func (m *TextMessage) AtAll() *TextMessage {
	m.At.IsAtAll = true
	return m
}

// AtUsers @指定用户
func (m *TextMessage) AtUsers(userIds ...string) *TextMessage {
	m.At.AtUserIds = append(m.At.AtUserIds, userIds...)
	return m
}

// ClearAtUsers 清除@的用户列表
func (m *TextMessage) ClearAtUsers() *TextMessage {
	m.At.AtUserIds = make([]string, 0)
	m.At.IsAtAll = false
	return m
}

// MarkdownMessage 表示markdown格式的消息
type MarkdownMessage struct {
	MsgType  string          `json:"msgtype"`
	Markdown MarkdownContent `json:"markdown"`
	At       MarkdownAt      `json:"at"`
}

// MarkdownContent 表示markdown消息内容
type MarkdownContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// MarkdownAt 表示markdown消息的@信息
type MarkdownAt struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

// NewMarkdownMessage 创建一个新的markdown消息
func NewMarkdownMessage() *MarkdownMessage {
	return &MarkdownMessage{
		MsgType:  "markdown",
		Markdown: MarkdownContent{},
		At: MarkdownAt{
			AtMobiles: make([]string, 0),
			AtUserIds: make([]string, 0),
		},
	}
}

// SetTitle 设置消息标题
func (m *MarkdownMessage) SetTitle(title string) *MarkdownMessage {
	m.Markdown.Title = title
	return m
}

// SetText 设置markdown格式的文本内容
func (m *MarkdownMessage) SetText(text string) *MarkdownMessage {
	m.Markdown.Text = text
	return m
}

// AtAll @所有人
func (m *MarkdownMessage) AtAll() *MarkdownMessage {
	m.At.IsAtAll = true
	return m
}

// AtMobiles @指定手机号的用户
func (m *MarkdownMessage) AtMobiles(mobiles ...string) *MarkdownMessage {
	m.At.AtMobiles = append(m.At.AtMobiles, mobiles...)
	return m
}

// AtUsers @指定用户ID的用户
func (m *MarkdownMessage) AtUsers(userIds ...string) *MarkdownMessage {
	m.At.AtUserIds = append(m.At.AtUserIds, userIds...)
	return m
}

// ClearAt 清除所有@信息
func (m *MarkdownMessage) ClearAt() *MarkdownMessage {
	m.At.AtMobiles = make([]string, 0)
	m.At.AtUserIds = make([]string, 0)
	m.At.IsAtAll = false
	return m
} 