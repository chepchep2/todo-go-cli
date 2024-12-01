package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	TasksFilePath string
	Environment   string
}

const (
	EnvTest = "test"
	EnvProd = "prod"
)

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
	return NewConfigWithEnv(EnvProd)
}

// NewTestConfig creates a new Config instance for testing
func NewTestConfig() (*Config, error) {
	return NewConfigWithEnv(EnvTest)
}

// NewConfigWithEnv creates a new Config instance with specified environment
func NewConfigWithEnv(env string) (*Config, error) {
	projectRoot, err := GetProjectRootPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get project root: %w", err)
	}

	var dataDir string
	if env == EnvTest {
		dataDir = filepath.Join(projectRoot, "testdataㅋ")
	} else {
		dataDir = filepath.Join(projectRoot, "data")
	}

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	return &Config{
		TasksFilePath: filepath.Join(dataDir, "tasks.json"),
		Environment:   env,
	}, nil
}
