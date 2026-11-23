package app

import (
	"errors"
	"strings"

	"github.com/abh2020-B2/todo-cli/internal/task"
)

// Todo manages a collection of tasks.
type Todo struct {
	tasks  []task.Task
	nextID int
}

// NewTodo creates and returns a new Todo instance.
func NewTodo() *Todo {
	return NewTodoWithTasks(nil)
}

// NewTodoWithTasks creates a Todo instance using existing tasks.
func NewTodoWithTasks(tasks []task.Task) *Todo {
	items := make([]task.Task, len(tasks))
	copy(items, tasks)

	todo := &Todo{
		tasks:  items,
		nextID: 1,
	}

	// Find the next available task ID.
	for _, item := range todo.tasks {
		if item.ID >= todo.nextID {
			todo.nextID = item.ID + 1
		}
	}

	return todo
}

// AddTask creates a new task and adds it to the task list.
func (t *Todo) AddTask(title string) (task.Task, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return task.Task{}, errors.New("task title cannot be empty")
	}

	newTask := task.Task{
		ID:        t.nextID,
		Title:     title,
		Completed: false,
	}

	t.tasks = append(t.tasks, newTask)
	t.nextID++

	return newTask, nil
}

// ListTasks returns all tasks stored in the Todo instance.
func (t *Todo) ListTasks() []task.Task {
	items := make([]task.Task, len(t.tasks))
	copy(items, t.tasks)

	return items
}

// CompleteTask marks a task as completed using its ID.
func (t *Todo) CompleteTask(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Completed = true
			return nil
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task using its ID.
func (t *Todo) DeleteTask(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

// EditTask changes the title of an existing task.
func (t *Todo) EditTask(id int, title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return errors.New("task title cannot be empty")
	}

	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Title = title
			return nil
		}
	}

	return errors.New("task not found")
}
