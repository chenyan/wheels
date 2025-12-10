package duckgo

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Rank       int    // 排名
	Title      string // 标题
	URL        string // 链接
	DisplayURL string // 显示的URL
	Source     string // 来源网站名称
	Snippet    string // 摘要
	Date       string // 日期（如果有）
	IsAd       bool   // 是否是广告
	IconURL    string // 图标URL
	HasDate    bool   // 是否有日期
}

func ParseNormalResult(html string) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("解析 HTML 失败: %v", err)
	}

	var results []SearchResult

	// 查找所有有机搜索结果
	doc.Find("ol.react-results--main li[data-layout='organic']").Each(func(i int, s *goquery.Selection) {
		result := SearchResult{
			Rank: i + 1,
		}

		// 提取标题
		titleElem := s.Find("h2.LnpumSThxEWMIsDdAT17 a.eVNpHGjtxRBq_gLOfGDr")
		result.Title = strings.TrimSpace(titleElem.Text())

		// 提取链接
		if href, exists := titleElem.Attr("href"); exists {
			result.URL = href
		}

		// 提取显示的URL
		displayURLElem := s.Find("a.Rn_JXVtoPVAFyGkcaXyK .veU5I0hFkgFGOPhX2RBE")
		result.DisplayURL = strings.TrimSpace(displayURLElem.Text())

		// 提取来源网站名称
		sourceElem := s.Find("p.fOCEb2mA3YZTJXXjpgdS")
		result.Source = strings.TrimSpace(sourceElem.Text())

		// 提取摘要
		snippetElem := s.Find("div.OgdwYG6KE2qthn9XQWFC span.kY2IgmnCmOGjharHErah")
		result.Snippet = strings.TrimSpace(snippetElem.Text())

		// 提取日期（如果有）
		dateElem := s.Find("span.MILR5XIVy9h75WrLvKiq")
		result.Date = strings.TrimSpace(dateElem.Text())

		// 只添加有标题的结果
		if result.Title != "" {
			results = append(results, result)
		}
	})

	return results, nil
}

func ParseStaticResult(html string) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("解析 HTML 失败: %v", err)
	}

	var results []SearchResult

	// 查找所有搜索结果（包括广告和普通结果）
	doc.Find("div.result.results_links.results_links_deep").Each(func(i int, s *goquery.Selection) {
		result := SearchResult{
			Rank: len(results) + 1,
		}

		// 检查是否是广告
		result.IsAd = s.HasClass("result--ad")

		// 提取标题
		titleElem := s.Find("h2.result__title a.result__a")
		result.Title = strings.TrimSpace(titleElem.Text())

		// 提取链接
		if href, exists := titleElem.Attr("href"); exists {
			// 解码DuckDuckGo的重定向链接
			result.URL = decodeDuckDuckGoURL(href)
		}

		// 提取显示的URL（可能有多种方式）
		// 方式1: 普通结果的显示URL
		displayURLElem := s.Find("div.result__extras__url a.result__url")
		if displayURLElem.Length() > 0 {
			result.DisplayURL = strings.TrimSpace(displayURLElem.Text())
		} else {
			// 方式2: 广告的特殊显示方式
			adURLElem := s.Find("a.result__url.sep--after")
			if adURLElem.Length() > 0 {
				result.DisplayURL = strings.TrimSpace(adURLElem.Text())
			} else {
				// 方式3: 其他位置的显示URL
				altURLElem := s.Find("a.result__url:not(.sep--after)")
				if altURLElem.Length() > 0 {
					result.DisplayURL = strings.TrimSpace(altURLElem.Text())
				}
			}
		}

		// 提取图标
		iconElem := s.Find("span.result__icon img.result__icon__img")
		if src, exists := iconElem.Attr("src"); exists {
			result.IconURL = src
		}

		// 提取摘要
		snippetElem := s.Find("a.result__snippet")
		result.Snippet = strings.TrimSpace(snippetElem.Text())

		// 提取日期
		dateText := ""
		dateSpan := s.Find("div.result__extras__url span")
		if dateSpan.Length() > 0 {
			// 找到包含日期的span（排除其他span）
			dateSpan.Each(func(j int, span *goquery.Selection) {
				text := strings.TrimSpace(span.Text())
				if strings.Contains(text, "T") && strings.Contains(text, "-") {
					dateText = text
				}
			})
		}

		if dateText != "" {
			// 解析日期，尝试多种格式
			dateText = strings.TrimSpace(dateText)

			// 移除时区信息以便解析
			if idx := strings.Index(dateText, "."); idx != -1 {
				dateText = dateText[:idx]
			}

			result.Date = dateText
		}

		// 只添加有标题的结果
		if result.Title != "" {
			results = append(results, result)
		}
	})

	return results, nil

}

// decodeDuckDuckGoURL 解码DuckDuckGo的重定向URL
func decodeDuckDuckGoURL(url string) string {
	// 这是一个简化的解码函数，实际可能需要更复杂的解析
	// DuckDuckGo的重定向URL格式：//duckduckgo.com/l/?uddg=原始URL编码...

	if strings.Contains(url, "uddg=") {
		// 提取uddg参数
		parts := strings.Split(url, "uddg=")
		if len(parts) > 1 {
			encodedURL := parts[1]

			// 可能需要提取到下一个参数之前的部分
			if ampIndex := strings.Index(encodedURL, "&"); ampIndex != -1 {
				encodedURL = encodedURL[:ampIndex]
			}

			// 简单的URL解码
			decodedURL, err := decodeURL(encodedURL)
			if err == nil {
				return decodedURL
			}
		}
	}

	return url
}

// decodeURL 简单的URL解码（简化版）
func decodeURL(encoded string) (string, error) {
	// 这里可以使用url.QueryUnescape，但需要先处理可能的编码问题
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return encoded, err
	}
	return decoded, nil
}
