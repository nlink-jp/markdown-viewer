package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_Defaults(t *testing.T) {
	// Clear env to avoid interference
	os.Unsetenv("MDV_PORT")
	os.Unsetenv("MDV_OPEN")
	os.Unsetenv("MDV_TARGET_DIR")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}
	if cfg.Port != 8080 {
		t.Errorf("Port = %d, want 8080", cfg.Port)
	}
	if cfg.Open != false {
		t.Error("Open should default to false")
	}
	if cfg.TargetDir != "." {
		t.Errorf("TargetDir = %q, want \".\"", cfg.TargetDir)
	}
}

func TestLoadConfig_EnvOverride(t *testing.T) {
	t.Setenv("MDV_PORT", "9999")
	t.Setenv("MDV_OPEN", "true")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}
	if cfg.Port != 9999 {
		t.Errorf("Port = %d, want 9999", cfg.Port)
	}
	if cfg.Open != true {
		t.Error("Open should be true from env")
	}
}

func TestLoadConfig_FileOverride(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.json")
	if err := os.WriteFile(cfgPath, []byte(`{"port": 3000, "open": true}`), 0o600); err != nil {
		t.Fatal(err)
	}

	// Change to temp dir so viper finds config.json
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	os.Unsetenv("MDV_PORT")
	os.Unsetenv("MDV_OPEN")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}
	if cfg.Port != 3000 {
		t.Errorf("Port = %d, want 3000", cfg.Port)
	}
}
