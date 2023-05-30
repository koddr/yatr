package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/koddr/gosl"
)

// runTasks provides the running tasks process.
func (q *Queue) runTasks() *Results {
	// Create a new slices for success and fail results.
	successResults := make([]Result, 0)
	failResults := make([]Result, 0)

	// Create a new wait groups for async and sequential queues.
	var wg sync.WaitGroup

	// Check, if async queue have tasks.
	if len(q.AsyncQueue) > 0 {
		// Loop for the all tasks in the async queue.
		for _, t := range q.AsyncQueue {
			// Loop for the all commands to execution in the single async task.
			for _, c := range t.Commands {
				// Add 1 to the wait group.
				wg.Add(1)

				// Create a new goroutine for the single command.
				go func(c Command) {
					// Don't forget to down the current wait group's count before goroutine is done.
					defer wg.Done()

					// Generate a new ID for current command.
					id, _ := gosl.RandomString(8)

					// Create a new struct for results.
					r := Result{ID: fmt.Sprintf("(async) %v", id), Name: c.Name, Description: c.Description}

					// Create a process for execute the current command.
					cmd := exec.Command(c.Exec[0], c.Exec[1:]...)

					// Check, if the current command want to run with sudo.
					if c.IsSudo {
						// Create a process for execute the current command as the sudo user.
						cmd = exec.Command("sudo", c.Exec...)
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
				}(c)
			}
		}
	}

	// Blocks all goroutines and wait the running commands.
	wg.Wait()

	// Check, if sequential queue have tasks.
	if len(q.SequentialQueue) > 0 {
		// Loop for the all tasks in the sequential queue.
		for _, t := range q.SequentialQueue {
			// Loop for the all commands to execution in the single sequential task.
			for _, c := range t.Commands {
				// Add 1 to the wait group.
				wg.Add(1)

				// Generate a new ID for current command.
				id, _ := gosl.RandomString(8)

				// Create a new struct for results.
				r := Result{ID: id, Name: c.Name, Description: c.Description}

				// Create a process for execute the current command.
				cmd := exec.Command(c.Exec[0], c.Exec[1:]...)

				// Check, if the current command want to run with sudo.
				if c.IsSudo {
					// Create a process for execute the current command as the sudo user.
					cmd = exec.Command("sudo", c.Exec...)
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
	}

	// Blocks all goroutines and wait the running commands.
	wg.Wait()

	return &Results{Success: successResults, Fail: failResults}
}
