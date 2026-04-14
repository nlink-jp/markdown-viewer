package browser

import (
	"testing"
)

func TestOpen_InvalidURL(t *testing.T) {
	err := Open("not-a-url")
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}

func TestOpen_NonHTTPScheme(t *testing.T) {
	tests := []string{
		"ftp://example.com",
		"javascript:alert(1)",
		"file:///etc/passwd",
		"data:text/html,<h1>test</h1>",
	}
	for _, url := range tests {
		t.Run(url, func(t *testing.T) {
			err := Open(url)
			if err == nil {
				t.Errorf("expected error for non-http scheme: %s", url)
			}
		})
	}
}

func TestOpen_EmptyURL(t *testing.T) {
	err := Open("")
	if err == nil {
		t.Error("expected error for empty URL")
	}
}
