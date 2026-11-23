package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abh2020-B2/todo-cli/internal/app"
	"github.com/abh2020-B2/todo-cli/internal/storage"
	"github.com/abh2020-B2/todo-cli/internal/task"
)

// Run processes CLI commands using the default storage path.
func Run(args []string) error {
	if len(args) == 0 || isHelp(args[0]) {
		printUsage()
		return nil
	}

	// Get the default storage path.
	path, err := storage.DefaultPath()
	if err != nil {
		return err
	}

	return runWithPath(args, path)
}

// runWithPath processes commands using a specified storage path.
func runWithPath(args []string, path string) error {
	if len(args) == 0 || isHelp(args[0]) {
		printUsage()
		return nil
	}

	// Load existing tasks.
	tasks, err := storage.Load(path)
	if err != nil {
		return err
	}

	// Create Todo using the loaded tasks.
	todo := app.NewTodoWithTasks(tasks)

	switch args[0] {

	case "add":
		if len(args) < 2 {
			return fmt.Errorf("please provide a task title")
		}

		title := strings.Join(args[1:], " ")

		newTask, err := todo.AddTask(title)
		if err != nil {
			return err
		}

		if err := saveTasks(path, todo); err != nil {
			return err
		}

		fmt.Printf("Task added successfully: [%d] %s\n",
			newTask.ID, newTask.Title)

	case "list":
		if len(args) > 2 {
			return fmt.Errorf("invalid list arguments")
		}

		filter := ""

		if len(args) == 2 {
			filter = args[1]

			if filter != "--completed" && filter != "--pending" {
				return fmt.Errorf("unknown list filter: %s", filter)
			}
		}

		displayTasks(todo.ListTasks(), filter)

	case "done":
		if len(args) != 2 {
			return fmt.Errorf("usage: todo done <id>")
		}

		id, err := parseID(args[1])
		if err != nil {
			return err
		}

		if err := todo.CompleteTask(id); err != nil {
			return err
		}

		if err := saveTasks(path, todo); err != nil {
			return err
		}

		fmt.Printf("Task %d marked as completed.\n", id)

	case "delete":
		if len(args) != 2 {
			return fmt.Errorf("usage: todo delete <id>")
		}

		id, err := parseID(args[1])
		if err != nil {
			return err
		}

		if err := todo.DeleteTask(id); err != nil {
			return err
		}

		if err := saveTasks(path, todo); err != nil {
			return err
		}

		fmt.Printf("Task %d deleted successfully.\n", id)

	case "edit":
		if len(args) < 3 {
			return fmt.Errorf("usage: todo edit <id> <new title>")
		}

		id, err := parseID(args[1])
		if err != nil {
			return err
		}

		title := strings.Join(args[2:], " ")

		if err := todo.EditTask(id, title); err != nil {
			return err
		}

		if err := saveTasks(path, todo); err != nil {
			return err
		}

		fmt.Printf("Task %d updated successfully.\n", id)

	default:
		return fmt.Errorf("unknown command: %s\nUse 'todo help' to see available commands", args[0])
	}

	return nil
}

// isHelp checks whether the user requested help.
func isHelp(command string) bool {
	return command == "help" ||
		command == "--help" ||
		command == "-h"
}

// saveTasks saves the current task list to JSON.
func saveTasks(path string, todo *app.Todo) error {
	return storage.Save(path, todo.ListTasks())
}

// parseID converts a string ID to an integer.
func parseID(value string) (int, error) {
	id, err := strconv.Atoi(value)

	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid task ID: %s", value)
	}

	return id, nil
}

// displayTasks prints tasks according to the selected filter.
func displayTasks(tasks []task.Task, filter string) {
	count := 0

	for _, item := range tasks {
		if filter == "--completed" && !item.Completed {
			continue
		}

		if filter == "--pending" && item.Completed {
			continue
		}

		status := " "

		if item.Completed {
			status = "x"
		}

		fmt.Printf("[%s] [%d] %s\n",
			status, item.ID, item.Title)

		count++
	}

	if count == 0 {
		fmt.Println("No tasks found.")
	}
}

// printUsage displays available commands.
func printUsage() {
	fmt.Println("TODO CLI - Task Manager")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  todo add <title>")
	fmt.Println("  todo list")
	fmt.Println("  todo list --completed")
	fmt.Println("  todo list --pending")
	fmt.Println("  todo done <id>")
	fmt.Println("  todo delete <id>")
	fmt.Println("  todo edit <id> <new title>")
	fmt.Println("  todo help")
}
