package main

import "errors"

// newQueue provides the queueing process.
func newQueue(tt *Tasks) (*Queue, error) {
	// Check, if tasks are existed and greater than zero.
	if tt.Tasks == nil || len(tt.Tasks) == 0 {
		return nil, errors.New("can't create a new runner without any tasks")
	}

	// Create slices for tasks queue.
	asyncTasks := make([]Task, 0)
	sequentialTasks := make([]Task, 0)

	// Set async and sequential tasks to separated slices.
	for _, t := range tt.Tasks {
		if t.IsAsync {
			// Add async task to slice.
			asyncTasks = append(asyncTasks, t)
		} else {
			// Add sequential task to slice.
			sequentialTasks = append(sequentialTasks, t)
		}
	}

	return &Queue{AsyncQueue: asyncTasks, SequentialQueue: sequentialTasks}, nil
}
