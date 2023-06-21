package main

// App represents a struct for the application instance.
type App struct {
	Tasks  *Tasks
	Queues *Queues
}

// Tasks represents a struct for tasks set.
type Tasks struct {
	Name        string  `koanf:"name"`
	Description string  `koanf:"description"`
	Data        []*Task `koanf:"tasks"`
}

// Task represents a struct for the single task.
type Task struct {
	IsAsync       bool     `koanf:"is_async"`
	IsSudo        bool     `koanf:"is_sudo"`
	IsPrintOutput bool     `koanf:"is_print_output"`
	Name          string   `koanf:"name"`
	Description   string   `koanf:"description"`
	Exec          []string `koanf:"exec"`
}

// Queues represents a struct for async and sequential tasks.
type Queues struct {
	Async, Sequential []*Task
}

// Result represents a struct for the single result.
type Result struct {
	ID, Name, Description, Output string
}
