package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	TasksFilePath string
}

// GetProjectRootPath returns the absolute path of the project root
func GetProjectRootPath() (string, error) {
	// Get the directory of the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current file path")
	}

	// Navigate up two directories from internal/config to project root
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	absProjectRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return absProjectRoot, nil
}

// NewConfig creates a new Config instance
func NewConfig() (*Config, error) {
	projectRoot, err := GetProjectRootPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get project root: %w", err)
	}

	dataDir := filepath.Join(projectRoot, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	return &Config{
		TasksFilePath: filepath.Join(dataDir, "tasks.json"),
	}, nil
}
