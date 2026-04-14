package server

import (
	"testing"
)

func TestContainsDotDot(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"/normal/path", false},
		{"/", false},
		{"", false},
		{"/view/readme.md", false},
		{"/../etc/passwd", true},
		{"/path/../secret", true},
		{"/..", true},
		{"../", true},
		{"/path/..hidden", false}, // ".." must be a full segment
		{"/path/...triple", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := containsDotDot(tt.path)
			if got != tt.want {
				t.Errorf("containsDotDot(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
