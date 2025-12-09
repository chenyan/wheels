package tavily

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	apiKey := os.Getenv("TAVILY_API_KEY")
	if apiKey == "" {
		t.Skip("TAVILY_API_KEY not set")
	}

	client := NewClient(apiKey)

	req := &SearchRequest{
		Query:         "who is Leo Messi?",
		Topic:         "general",
		SearchDepth:   "basic",
		MaxResults:    3,
		IncludeAnswer: true,
	}

	resp, err := client.Search(req)
	if err != nil {
		t.Fatalf("search error: %v", err)
	}

	t.Logf("Query: %s", resp.Query)
	t.Logf("Answer: %s", resp.Answer)
	t.Logf("ResponseTime: %v", resp.ResponseTime)
	t.Logf("RequestID: %s", resp.RequestID)

	for i, result := range resp.Results {
		t.Logf("Result %d:", i+1)
		t.Logf("  Title: %s", result.Title)
		t.Logf("  URL: %s", result.URL)
		t.Logf("  Score: %.4f", result.Score)
		t.Logf("  Content: %s", result.Content[:min(100, len(result.Content))]+"...")
	}
}

func TestQuickSearch(t *testing.T) {
	apiKey := os.Getenv("TAVILY_API_KEY")
	if apiKey == "" {
		t.Skip("TAVILY_API_KEY not set")
	}

	client := NewClient(apiKey)
	resp, err := client.QuickSearch("latest news about AI", 5)
	if err != nil {
		t.Fatalf("quick search error: %v", err)
	}

	t.Logf("Found %d results", len(resp.Results))
}
