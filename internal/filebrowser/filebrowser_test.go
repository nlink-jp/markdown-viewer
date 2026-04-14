package filebrowser

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Create directories
	os.MkdirAll(filepath.Join(dir, "subdir"), 0o755)
	os.MkdirAll(filepath.Join(dir, "another"), 0o755)

	// Create markdown files
	os.WriteFile(filepath.Join(dir, "readme.md"), []byte("# Hello"), 0o644)
	os.WriteFile(filepath.Join(dir, "guide.markdown"), []byte("# Guide"), 0o644)

	// Create non-markdown files (should be excluded)
	os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("text"), 0o644)
	os.WriteFile(filepath.Join(dir, "data.json"), []byte("{}"), 0o644)

	// Create markdown file in subdir
	os.WriteFile(filepath.Join(dir, "subdir", "nested.md"), []byte("# Nested"), 0o644)

	return dir
}

func TestListDirectory_Root(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	// Expect: 2 dirs (another, subdir) + 2 md files (guide.markdown, readme.md)
	if len(items) != 4 {
		t.Errorf("got %d items, want 4", len(items))
		for _, item := range items {
			t.Logf("  %s (dir=%v)", item.Name, item.IsDir)
		}
	}
}

func TestListDirectory_DirsFirst(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	if len(items) < 3 {
		t.Fatal("not enough items")
	}

	// Directories should come before files
	if !items[0].IsDir {
		t.Error("first item should be a directory")
	}
	// Last items should be files
	if items[len(items)-1].IsDir {
		t.Error("last item should be a file")
	}
}

func TestListDirectory_ExcludesNonMarkdown(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	for _, item := range items {
		if !item.IsDir && item.Name == "notes.txt" {
			t.Error("non-markdown file should be excluded: notes.txt")
		}
		if !item.IsDir && item.Name == "data.json" {
			t.Error("non-markdown file should be excluded: data.json")
		}
	}
}

func TestListDirectory_MarkdownExtensions(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	fileNames := map[string]bool{}
	for _, item := range items {
		if !item.IsDir {
			fileNames[item.Name] = true
		}
	}

	if !fileNames["readme.md"] {
		t.Error("should include .md files")
	}
	if !fileNames["guide.markdown"] {
		t.Error("should include .markdown files")
	}
}

func TestListDirectory_Subdirectory(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/subdir")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("got %d items, want 1", len(items))
	}
	if items[0].Name != "nested.md" {
		t.Errorf("name = %q, want nested.md", items[0].Name)
	}
}

func TestListDirectory_PathFormat(t *testing.T) {
	dir := setupTestDir(t)

	items, err := ListDirectory(dir, "/subdir")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	if len(items) != 1 {
		t.Fatal("expected 1 item")
	}
	// Path should use forward slashes
	if items[0].Path != "/subdir/nested.md" {
		t.Errorf("path = %q, want /subdir/nested.md", items[0].Path)
	}
}

func TestListDirectory_NonExistent(t *testing.T) {
	_, err := ListDirectory("/nonexistent", "/")
	if err == nil {
		t.Error("expected error for non-existent directory")
	}
}

func TestListDirectory_Empty(t *testing.T) {
	dir := t.TempDir()

	items, err := ListDirectory(dir, "/")
	if err != nil {
		t.Fatalf("ListDirectory: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("got %d items, want 0", len(items))
	}
}
