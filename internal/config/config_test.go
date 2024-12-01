package config_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo-go-cli/internal/config"
)

var _ = Describe("Config", func() {
	var cfg *config.Config

	BeforeEach(func() {
		var err error
		cfg, err = config.NewTestConfig()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// Clean up test data
		projectRoot, err := config.GetProjectRootPath()
		Expect(err).NotTo(HaveOccurred())
		testDataDir := filepath.Join(projectRoot, "testdata")
		_ = os.RemoveAll(testDataDir)
	})

	It("should create config with test environment", func() {
		Expect(cfg.Environment).To(Equal(config.EnvTest))
		Expect(cfg.TasksFilePath).To(ContainSubstring("testdata"))
	})

	It("should create data directory if it doesn't exist", func() {
		projectRoot, err := config.GetProjectRootPath()
		Expect(err).NotTo(HaveOccurred())
		testDataDir := filepath.Join(projectRoot, "testdata")
		
		// Directory should exist
		_, err = os.Stat(testDataDir)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when creating production config", func() {
		It("should create config with production environment", func() {
			prodCfg, err := config.NewConfig()
			Expect(err).NotTo(HaveOccurred())
			Expect(prodCfg.Environment).To(Equal(config.EnvProd))
			Expect(prodCfg.TasksFilePath).To(ContainSubstring("data"))
		})
	})
})
