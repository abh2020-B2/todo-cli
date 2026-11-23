package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abh2020-B2/todo-cli/internal/task"
)

// DefaultPath returns the default path for the tasks JSON file.
func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get user home directory: %w", err)
	}

	return filepath.Join(home, ".todo-cli", "tasks.json"), nil
}

// Load reads tasks from a JSON file.
func Load(path string) ([]task.Task, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		// Return an empty task list if the file does not exist.
		if errors.Is(err, os.ErrNotExist) {
			return []task.Task{}, nil
		}

		return nil, fmt.Errorf("read tasks file: %w", err)
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("decode tasks JSON: %w", err)
	}

	// Ensure that the returned slice is not nil.
	if tasks == nil {
		tasks = []task.Task{}
	}

	return tasks, nil
}

// Save writes tasks to a JSON file.
func Save(path string, tasks []task.Task) error {
	// Convert tasks to formatted JSON.
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("encode tasks JSON: %w", err)
	}

	// Create the directory if it does not exist.
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create storage directory: %w", err)
	}

	// Add a newline at the end of the JSON file.
	data = append(data, '\n')

	// Write the JSON data to the file.
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}

	return nil
}
