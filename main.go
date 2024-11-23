package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID   int    `json: "id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var tasks []Task

func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tasks)
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)
}

func main() {

	if err := loadTasks(); err != nil {
		tasks = []Task{}
	}

	// 인자가 없으면 사용법 출력
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	args := os.Args[1:]
	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("할일을 입력하세요")
			return
		}
		newTask := Task{
			ID:   len(tasks) + 1,
			Text: args[1],
			Done: false,
		}
		tasks = append(tasks, newTask)

		if err := saveTasks(); err != nil {
			fmt.Printf("저장 실패: %v\n", err)
			return
		}

		fmt.Printf("Task added: %s\n", args[1])
	case "list":
		if len(tasks) == 0 {
			fmt.Println("할일이 없습니다")
			return
		}
		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "x"
			}

			fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
		}
	case "done":
		if len(args) < 2 {
			fmt.Println("list 번호를 입력하세요.")
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("유효한 번호를 입력하세요.")
			return
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("할일 %d번을 찾을 수 없습니다.\n", id)
			return
		}

		if err := saveTasks(); err != nil {
			fmt.Printf("저장 실패: %v\n", err)
			return
		}

		fmt.Printf("할일 %d번이 완료되었습니다.\n", id)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add [task]    Add a new task")
	fmt.Println("  todo list          Show all tasks")
	fmt.Println("  todo done [id]     Mark a task as done")
}
