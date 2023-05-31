package main

import (
	"errors"
	"fmt"
)

// newQueue provides the queueing process.
func newQueue(tt *Tasks) (*Queue, error) {
	// Check, if tasks are existed and greater than zero.
	if tt.Tasks == nil || len(tt.Tasks) == 0 {
		return nil, errors.New("error: can't create a new queue without any tasks")
	}

	// Create slices for tasks queue.
	asyncTasks := make([]Task, 0)
	sequentialTasks := make([]Task, 0)

	// Set async and sequential tasks to separated slices.
	for i, t := range tt.Tasks {
		// Check, if the current task's command is set.
		if t.Name == "" {
			// Return error with the current task index.
			return nil, fmt.Errorf("error: task %d has no name", i+1)
			// Check, if the current task's command is set.
		} else if t.Exec == nil || len(t.Exec) == 0 {
			// Return error with the current task name.
			return nil, fmt.Errorf("error: task %d (name: '%s') has no commands to execute", i+1, t.Name)
		} else {
			// Check, if the current task is an async.
			if t.IsAsync {
				// Add async task to slice.
				asyncTasks = append(asyncTasks, t)
			} else {
				// Add sequential task to slice.
				sequentialTasks = append(sequentialTasks, t)
			}
		}
	}

	return &Queue{AsyncQueue: asyncTasks, SequentialQueue: sequentialTasks}, nil
}
