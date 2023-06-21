package main

import (
	"fmt"

	"github.com/koddr/gosl"
)

// newTasks provides a new Tasks instance.
func newTasks() (*Tasks, error) {
	// Check, if the file name is too short.
	if pathFlag == "" {
		return nil, fmt.Errorf("invalid format of tasks file, see: %s", WikiPageURL)
	}

	// Create a new config instance.
	tasks := &Tasks{}

	// Load config from path or HTTP by the given file format.
	_, err := gosl.ParseFileWithEnvToStruct(pathFlag, "YATR", tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
