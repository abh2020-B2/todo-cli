package cli

import (
	"path/filepath"
	"testing"

	"github.com/abh2020-B2/todo-cli/internal/storage"
)

// TestRunAddTask checks whether the add command saves a task.
func TestRunAddTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	err := runWithPath([]string{"add", "Learn", "Go"}, path)
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := storage.Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Learn Go" {
		t.Errorf("unexpected title: %s", tasks[0].Title)
	}
}

// TestRunTaskLifecycle checks add, complete, edit and delete commands.
func TestRunTaskLifecycle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	// Add a task.
	if err := runWithPath([]string{"add", "Learn Go"}, path); err != nil {
		t.Fatal(err)
	}

	// Complete the task.
	if err := runWithPath([]string{"done", "1"}, path); err != nil {
		t.Fatal(err)
	}

	tasks, err := storage.Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if !tasks[0].Completed {
		t.Error("expected task to be completed")
	}

	// Edit the task.
	if err := runWithPath(
		[]string{"edit", "1", "Learn Advanced Go"}, path,
	); err != nil {
		t.Fatal(err)
	}

	tasks, err = storage.Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if tasks[0].Title != "Learn Advanced Go" {
		t.Errorf("unexpected title: %s", tasks[0].Title)
	}

	// Delete the task.
	if err := runWithPath([]string{"delete", "1"}, path); err != nil {
		t.Fatal(err)
	}

	tasks, err = storage.Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}

// TestRunInvalidCommands checks invalid CLI arguments.
func TestRunInvalidCommands(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "missing task title",
			args: []string{"add"},
		},
		{
			name: "invalid task ID",
			args: []string{"done", "abc"},
		},
		{
			name: "unknown command",
			args: []string{"unknown"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "tasks.json")

			err := runWithPath(test.args, path)

			if err == nil {
				t.Error("expected an error, got nil")
			}
		})
	}
}
