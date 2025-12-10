package duckgo

import (
	"errors"
	"fmt"
	"slices"

	"github.com/chenyan/wheels/external/cloudflare/browser"
	"github.com/chenyan/wheels/funcs/seqs"
)

const (
	BaseURL = "https://html.duckduckgo.com/html/?q=%s"
)

var (
	cf *browser.Client
)

func Init() {
	cf = browser.NewClientWithEnv()
}

func Query(query string) ([]SearchResult, error) {
	if cf == nil || query == "" {
		return nil, errors.New("cf or query is empty")
	}

	url := fmt.Sprintf(BaseURL, query)
	html, err := cf.GetBrowserHTML(url)
	if err != nil {
		return nil, err
	}
	rs, err := ParseStaticResult(html)
	if err != nil {
		return nil, err
	}
	return slices.Collect(seqs.Filter(rs, func(r SearchResult) bool {
		return !r.IsAd
	})), nil
}
