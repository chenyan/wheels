package tavily

import (
	"fmt"
	"time"

	"github.com/chenyan/wheels/httpx/reqs"
)

const (
	BaseURL = "https://api.tavily.com"
)

// Client Tavily API 客户端
type Client struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
}

// NewClient 创建一个新的 Tavily API 客户端
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: BaseURL,
		Timeout: 30 * time.Second,
	}
}

// SetBaseURL 设置基础 URL
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.BaseURL = baseURL
	return c
}

// SetTimeout 设置超时时间
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.Timeout = timeout
	return c
}

// SearchRequest 搜索请求参数
type SearchRequest struct {
	Query                    string   `json:"query"`
	AutoParameters           bool     `json:"auto_parameters,omitempty"`
	Topic                    string   `json:"topic,omitempty"`
	SearchDepth              string   `json:"search_depth,omitempty"`
	ChunksPerSource          int      `json:"chunks_per_source,omitempty"`
	MaxResults               int      `json:"max_results,omitempty"`
	TimeRange                *string  `json:"time_range,omitempty"`
	StartDate                string   `json:"start_date,omitempty"`
	EndDate                  string   `json:"end_date,omitempty"`
	IncludeAnswer            bool     `json:"include_answer,omitempty"`
	IncludeRawContent        bool     `json:"include_raw_content,omitempty"`
	IncludeImages            bool     `json:"include_images,omitempty"`
	IncludeImageDescriptions bool     `json:"include_image_descriptions,omitempty"`
	IncludeFavicon           bool     `json:"include_favicon,omitempty"`
	IncludeDomains           []string `json:"include_domains,omitempty"`
	ExcludeDomains           []string `json:"exclude_domains,omitempty"`
	Country                  *string  `json:"country,omitempty"`
}

// SearchResult 单个搜索结果
type SearchResult struct {
	Title      string  `json:"title"`
	URL        string  `json:"url"`
	Content    string  `json:"content"`
	Score      float64 `json:"score"`
	RawContent *string `json:"raw_content"`
	Favicon    string  `json:"favicon,omitempty"`
}

// AutoParameters 自动参数
type AutoParameters struct {
	Topic       string `json:"topic"`
	SearchDepth string `json:"search_depth"`
}

// SearchResponse 搜索响应
type SearchResponse struct {
	Query          string          `json:"query"`
	Answer         string          `json:"answer,omitempty"`
	Images         []string        `json:"images,omitempty"`
	Results        []SearchResult  `json:"results"`
	ResponseTime   float32         `json:"response_time"`
	AutoParameters *AutoParameters `json:"auto_parameters,omitempty"`
	RequestID      string          `json:"request_id"`
}

// Search 执行搜索
func (c *Client) Search(req *SearchRequest) (*SearchResponse, error) {
	url := c.BaseURL + "/search"

	opts := &reqs.Opts{
		Headers: map[string]string{
			"Authorization": "Bearer " + c.APIKey,
		},
		Timeout: c.Timeout,
	}

	resp := reqs.PostJSON2(url, req, opts)
	if resp.Error != nil {
		return nil, resp.Error
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		body, _ := resp.String()
		return nil, fmt.Errorf("tavily api error: status=%d, body=%s", resp.StatusCode, body)
	}

	var result SearchResponse
	if err := resp.JSON(&result); err != nil {
		return nil, fmt.Errorf("decode response error: %w", err)
	}

	return &result, nil
}

// QuickSearch 快速搜索，使用默认参数
func (c *Client) QuickSearch(query string, maxResults int) (*SearchResponse, error) {
	req := &SearchRequest{
		Query:      query,
		MaxResults: maxResults,
	}
	return c.Search(req)
}
