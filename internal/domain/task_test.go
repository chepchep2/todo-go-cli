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

	Describe("MarkAsDone", func() {
		It("should mark the task as done", func() {
			Expect(task.Done).To(BeFalse())
			
			task.MarkAsDone()
			
			Expect(task.Done).To(BeTrue())
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
				task.MarkAsDone()
			})

			It("should return the correct string representation", func() {
				expected := "1. [x] Test task"
				Expect(task.String()).To(Equal(expected))
			})
		})
	})
})
