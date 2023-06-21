package main

import (
	"fmt"
)

// newQueues provides the queueing process.
func newQueues(tasks *Tasks) (*Queues, error) {
	// Check, if task set name is existed.
	if tasks.Name == "" {
		return nil, fmt.Errorf("can't create a new queue without a task set name, see: %s", WikiPageURL)
	}

	// Check, if tasks are existed and greater than zero.
	if tasks.Data == nil || len(tasks.Data) == 0 {
		return nil, fmt.Errorf("can't create a new queue without any tasks, see: %s", WikiPageURL)
	}

	// Create slices for tasks queue.
	asyncTasks := make([]*Task, 0)
	sequentialTasks := make([]*Task, 0)

	// Set async and sequential tasks to separated slices.
	for i, t := range tasks.Data {
		// Check, if the current task's command name is set.
		if t.Name == "" {
			// Return error with the current task index.
			return nil, fmt.Errorf("task %d has no name, see: %s", i+1, WikiPageURL)
		}

		// Check, if the current task's command to execute is set.
		if t.Exec == nil || len(t.Exec) == 0 {
			// Return error with the current task name.
			return nil, fmt.Errorf(
				"task %d (name: '%s') has no commands to execute, see: %s",
				i+1, t.Name, WikiPageURL,
			)
		}

		// Check, if the current task is an async.
		if t.IsAsync {
			// Add async task to slice.
			asyncTasks = append(asyncTasks, t)
		} else {
			// Add sequential task to slice.
			sequentialTasks = append(sequentialTasks, t)
		}
	}

	return &Queues{Async: asyncTasks, Sequential: sequentialTasks}, nil
}
