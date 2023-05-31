package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/koddr/gosl"
)

// runTasks provides the running tasks process.
func (q *Queue) runTasks() (*Results, error) {
	// Check, if success and sequential queue have tasks.
	if q.AsyncQueue == nil && q.SequentialQueue == nil {
		return nil, fmt.Errorf(
			"error: not possible to run tasks without any tasks, see: %s",
			WikiPageURL,
		)
	}

	// Create a new slices for success and fail results.
	successResults := make([]Result, 0)
	failResults := make([]Result, 0)

	// Create a new wait groups for async and sequential queues.
	var wg sync.WaitGroup

	// Check, if async queue have tasks.
	if len(q.AsyncQueue) > 0 {
		// Loop for the all tasks in the async queue.
		for _, t := range q.AsyncQueue {
			// Add 1 to the wait group.
			wg.Add(1)

			// Create a new goroutine for the single command.
			go func(t Task) {
				// Don't forget to down the current wait group's count before goroutine is done.
				defer wg.Done()

				// Generate a new ID for current command.
				id, _ := gosl.RandomString(8)

				// Create a new struct for results.
				r := Result{ID: fmt.Sprintf("(async) %v", id), Name: t.Name, Description: t.Description}

				// Create a process for execute the current command.
				cmd := exec.Command(t.Exec[0], t.Exec[1:]...)

				// Check, if the current command want to run with sudo.
				if t.IsSudo {
					// Create a process for execute the current command as the sudo user.
					cmd = exec.Command("sudo", t.Exec...)
				}

				// Run execution of the current command.
				if err := cmd.Run(); err != nil {
					// If error was happened, add record with stderr output to fail results.
					r.Output = err.Error()
					// Add record with an error message to fail results.
					failResults = append(failResults, r)
				} else {
					// Add record with command info to success results.
					successResults = append(successResults, r)
				}
			}(t)
		}
	}

	// Blocks all goroutines and wait the running commands.
	wg.Wait()

	// Check, if sequential queue have tasks.
	if len(q.SequentialQueue) > 0 {
		// Loop for the all tasks in the sequential queue.
		for _, t := range q.SequentialQueue {
			// Add 1 to the wait group.
			wg.Add(1)

			// Generate a new ID for current command.
			id, _ := gosl.RandomString(8)

			// Create a new struct for results.
			r := Result{ID: id, Name: t.Name, Description: t.Description}

			// Create a process for execute the current command.
			cmd := exec.Command(t.Exec[0], t.Exec[1:]...)

			// Check, if the current command want to run with sudo.
			if t.IsSudo {
				// Create a process for execute the current command as the sudo user.
				cmd = exec.Command("sudo", t.Exec...)
			}

			// Run execution of the current command.
			if err := cmd.Run(); err != nil {
				// If error was happened, add record with stderr output to fail results.
				r.Output = err.Error()
				// Add record with an error message to fail results.
				failResults = append(failResults, r)
			} else {
				// Add record with command info to success results.
				successResults = append(successResults, r)
			}

			// Don't forget to down the current wait group's count before goroutine is done.
			wg.Done()
		}
	}

	// Blocks all goroutines and wait the running commands.
	wg.Wait()

	return &Results{Success: successResults, Fail: failResults}, nil
}
