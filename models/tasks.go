package models

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTodos return todolist
func GetTodos() (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "todo1"},
			{2, "todo2"},
		},
	}

	return
}
