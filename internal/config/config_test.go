package config_test

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo-go-cli/internal/config"
)

var _ = Describe("Config", func() {
	Describe("GetProjectRootPath", func() {
		It("should return the project root path", func() {
			rootPath, err := config.GetProjectRootPath()
			Expect(err).NotTo(HaveOccurred())
			Expect(rootPath).To(ContainSubstring("todo-go-cli"))
			
			// Verify that go.mod exists in the root path
			_, err = os.Stat(filepath.Join(rootPath, "go.mod"))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("NewConfig", func() {
		var cfg *config.Config
		var err error

		BeforeEach(func() {
			cfg, err = config.NewConfig()
		})

		It("should create a new config without error", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).NotTo(BeNil())
		})

		It("should set tasks file path in the data directory", func() {
			Expect(cfg.TasksFilePath).To(ContainSubstring("data"))
			Expect(cfg.TasksFilePath).To(HaveSuffix("tasks.json"))
		})

		It("should create the data directory", func() {
			dataDir := filepath.Dir(cfg.TasksFilePath)
			info, err := os.Stat(dataDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeTrue())
		})

		AfterEach(func() {
			// Clean up the data directory after tests
			if cfg != nil && cfg.TasksFilePath != "" {
				dataDir := filepath.Dir(cfg.TasksFilePath)
				if strings.HasSuffix(dataDir, "data") {
					os.RemoveAll(dataDir)
				}
			}
		})
	})
})
