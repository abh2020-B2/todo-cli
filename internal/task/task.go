package task

// Task represents a single todo item.
type Task struct {
	ID        int    `json:"id"`        // Unique identifier of the task
	Title     string `json:"title"`     // Title or description of the task
	Completed bool   `json:"completed"` // Indicates whether the task is completed
}
