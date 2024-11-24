package domain

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// NewTask creates a new Task instance
func NewTask(id int, text string) *Task {
	return &Task{
		ID:   id,
		Text: text,
		Done: false,
	}
}

// MarkAsDone marks the task as done
func (t *Task) MarkAsDone() {
	t.Done = true
}

// String returns a string representation of the task
func (t *Task) String() string {
	status := " "
	if t.Done {
		status = "x"
	}
	return fmt.Sprintf("%d. [%s] %s", t.ID, status, t.Text)
}
