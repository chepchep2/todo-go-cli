package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo-go-cli/internal/domain"
)

var _ = Describe("Task", func() {
	var task *domain.Task

	BeforeEach(func() {
		task = domain.NewTask(1, "Test task")
	})

	Describe("NewTask", func() {
		It("should create a new task with the given ID and text", func() {
			Expect(task.ID).To(Equal(1))
			Expect(task.Text).To(Equal("Test task"))
			Expect(task.Done).To(BeFalse())
		})
	})

	Describe("ToggleAsDone", func() {
		It("should toggle the task's done status", func() {
			Expect(task.Done).To(BeFalse())

			task.ToggleAsDone()
			Expect(task.Done).To(BeTrue())

			task.ToggleAsDone()
			Expect(task.Done).To(BeFalse())
		})
	})

	Describe("String", func() {
		Context("when task is not done", func() {
			It("should return the correct string representation", func() {
				expected := "1. [ ] Test task"
				Expect(task.String()).To(Equal(expected))
			})
		})

		Context("when task is done", func() {
			BeforeEach(func() {
				task.ToggleAsDone()
			})

			It("should return the correct string representation", func() {
				expected := "1. [x] Test task"
				Expect(task.String()).To(Equal(expected))
			})
		})
	})
})
