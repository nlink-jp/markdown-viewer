package markdown

import (
	"testing"
)

func TestIsSafeLink(t *testing.T) {
	tests := []struct {
		name string
		dest string
		want bool
	}{
		// Safe: local markdown files
		{"md file", "readme.md", true},
		{"markdown file", "guide.markdown", true},
		{"relative path md", "./docs/readme.md", true},
		{"nested path md", "subdir/guide.markdown", true},
		{"uppercase MD", "README.MD", true},
		{"mixed case", "Guide.Markdown", true},

		// Unsafe: external links
		{"http", "http://example.com", false},
		{"https", "https://example.com", false},
		{"http with path", "http://example.com/page.md", false},

		// Unsafe: non-markdown local files
		{"txt file", "notes.txt", false},
		{"html file", "index.html", false},
		{"go file", "main.go", false},
		{"no extension", "README", false},
		{"json file", "data.json", false},
		{"empty", "", false},

		// Unsafe: protocol handlers
		{"javascript", "javascript:alert(1)", false},
		{"data uri", "data:text/html,<h1>test</h1>", false},
		{"ftp", "ftp://server/file.md", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSafeLink([]byte(tt.dest))
			if got != tt.want {
				t.Errorf("isSafeLink(%q) = %v, want %v", tt.dest, got, tt.want)
			}
		})
	}
}
