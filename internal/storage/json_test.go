package storage

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/abh2020-B2/todo-cli/internal/task"
)

// TestSaveAndLoad checks JSON saving and loading.
func TestSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")

	original := []task.Task{
		{ID: 1, Title: "Learn Go", Completed: false},
		{ID: 2, Title: "Build a project", Completed: true},
	}

	// Save tasks.
	if err := Save(path, original); err != nil {
		t.Fatal(err)
	}

	// Load tasks.
	loaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	// Compare original and loaded tasks.
	if !reflect.DeepEqual(original, loaded) {
		t.Errorf("loaded tasks do not match the original tasks")
	}
}

// TestLoadMissingFile checks loading a nonexistent file.
func TestLoadMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.json")

	tasks, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		t.Errorf("expected an empty task list, got %d tasks", len(tasks))
	}
}
