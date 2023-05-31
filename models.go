package main

// Queue presents a struct for the queue for async and sequential tasks.
type Queue struct {
	AsyncQueue, SequentialQueue []Task
}

// Results present a struct for success and fail results.
type Results struct {
	Success, Fail []Result
}

// Result presents struct for the single result.
type Result struct {
	ID, Name, Description, Output string
}

// Tasks present a struct for tasks set.
type Tasks struct {
	Name        string `koanf:"name"`
	Description string `koanf:"description"`
	Tasks       []Task `koanf:"tasks"`
}

// Task presents a struct for the single task.
type Task struct {
	IsAsync     bool     `koanf:"is_async"`
	IsSudo      bool     `koanf:"is_sudo"`
	Name        string   `koanf:"name"`
	Description string   `koanf:"description"`
	Exec        []string `koanf:"exec"`
}
