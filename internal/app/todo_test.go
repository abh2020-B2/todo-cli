package app

import (
	"testing"

	"github.com/abh2020-B2/todo-cli/internal/task"
)

// TestAddTask checks whether a new task is created.
func TestAddTask(t *testing.T) {
	todo := NewTodo()

	newTask, err := todo.AddTask("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	if newTask.ID != 1 {
		t.Errorf("expected ID 1, got %d", newTask.ID)
	}

	if newTask.Title != "Learn Go" {
		t.Errorf("unexpected title: %s", newTask.Title)
	}
}

// TestAddTaskEmptyTitle checks empty title validation.
func TestAddTaskEmptyTitle(t *testing.T) {
	todo := NewTodo()

	_, err := todo.AddTask("   ")
	if err == nil {
		t.Fatal("expected an error for an empty title")
	}
}

// TestCompleteTask checks task completion.
func TestCompleteTask(t *testing.T) {
	todo := NewTodo()

	newTask, err := todo.AddTask("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	err = todo.CompleteTask(newTask.ID)
	if err != nil {
		t.Fatal(err)
	}

	tasks := todo.ListTasks()

	if !tasks[0].Completed {
		t.Error("expected task to be completed")
	}
}

// TestCompleteTaskNotFound checks invalid task IDs.
func TestCompleteTaskNotFound(t *testing.T) {
	todo := NewTodo()

	err := todo.CompleteTask(100)
	if err == nil {
		t.Error("expected an error for a missing task")
	}
}

// TestDeleteTask checks task deletion.
func TestDeleteTask(t *testing.T) {
	todo := NewTodo()

	first, _ := todo.AddTask("Learn Go")
	_, _ = todo.AddTask("Build a project")

	err := todo.DeleteTask(first.ID)
	if err != nil {
		t.Fatal(err)
	}

	tasks := todo.ListTasks()

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Build a project" {
		t.Errorf("unexpected remaining task: %s", tasks[0].Title)
	}
}

// TestDeleteTaskNotFound checks deletion of a missing task.
func TestDeleteTaskNotFound(t *testing.T) {
	todo := NewTodo()

	err := todo.DeleteTask(100)
	if err == nil {
		t.Error("expected an error for a missing task")
	}
}

// TestEditTask checks task editing.
func TestEditTask(t *testing.T) {
	todo := NewTodo()

	newTask, err := todo.AddTask("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	err = todo.EditTask(newTask.ID, "Learn Advanced Go")
	if err != nil {
		t.Fatal(err)
	}

	tasks := todo.ListTasks()

	if tasks[0].Title != "Learn Advanced Go" {
		t.Errorf("unexpected title: %s", tasks[0].Title)
	}
}

// TestEditTaskEmptyTitle checks empty title validation.
func TestEditTaskEmptyTitle(t *testing.T) {
	todo := NewTodo()

	newTask, err := todo.AddTask("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	err = todo.EditTask(newTask.ID, "   ")
	if err == nil {
		t.Error("expected an error for an empty title")
	}
}

// TestNextIDAfterLoading checks ID generation after loading tasks.
func TestNextIDAfterLoading(t *testing.T) {
	tasks := []task.Task{
		{ID: 1, Title: "Task One"},
		{ID: 5, Title: "Task Five"},
	}

	todo := NewTodoWithTasks(tasks)

	newTask, err := todo.AddTask("Task Six")
	if err != nil {
		t.Fatal(err)
	}

	if newTask.ID != 6 {
		t.Errorf("expected ID 6, got %d", newTask.ID)
	}
}
