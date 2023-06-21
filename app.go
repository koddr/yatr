package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/koddr/gosl"
)

// newApp provides a new application instance.
func newApp(tasks *Tasks, queues *Queues) *App {
	return &App{
		Tasks:  tasks,
		Queues: queues,
	}
}

// start starts an application with a beauty output to console.
func (app *App) start() error {
	// Start timer.
	start := time.Now()

	printStyled(
		"Hello and welcome to yatr! 👋",
		"",
		"margin-top-bottom",
	)
	printStyled(
		fmt.Sprintf(
			"Run task set [name: %s, description: %s]",
			app.Tasks.Name, app.Tasks.Description,
		),
		"info",
		"",
	)

	// Check, if tasks data is not nil.
	if app.Tasks.Data != nil {
		// Check, if tasks data is not empty.
		if len(app.Tasks.Data) > 0 {
			// Create a new wait groups for async and sequential queues.
			var wg sync.WaitGroup

			// Create a new slices for success and fail results.
			successResults, failResults := make([]*Result, 0), make([]*Result, 0)

			printStyled(
				fmt.Sprintf(
					"According to tasks file in '%s', %d commands to execute (async: %d, sequential: %d).",
					pathFlag, len(app.Tasks.Data), len(app.Queues.Async), len(app.Queues.Sequential),
				),
				"info",
				"",
			)
			printStyled(
				"OK! Your commands will be launched now... Please wait!",
				"info",
				"",
			)

			// Check, if async queue have tasks.
			if len(app.Queues.Async) > 0 {
				// Loop for the all tasks in the async queue.
				for _, t := range app.Queues.Async {
					// Check, if current task have commands to execute.
					if t.Exec == nil || len(t.Exec) == 0 {
						return fmt.Errorf("can't run task without any commands to execute, see: %s", WikiPageURL)
					}

					// Add 1 to the wait group.
					wg.Add(1)

					// Create a new goroutine for the single command.
					go func(t *Task) {
						// Don't forget to down the current wait group's count before goroutine is done.
						defer wg.Done()

						// Create a new buffer for Stdout.
						var stdOut bytes.Buffer

						// Generate a new ID for current command.
						id, _ := gosl.RandomString(8)

						// Create a new struct for results.
						r := &Result{
							ID:          gosl.Concat("(async) ", id),
							Name:        t.Name,
							Description: t.Description,
						}

						// Create a process for execute the current command.
						cmd := exec.Command(t.Exec[0], t.Exec[1:]...)

						// Check, if the current command want to run with sudo.
						if t.IsSudo {
							// Create a process for execute the current command as the sudo user.
							cmd = exec.Command("sudo", t.Exec...)
						}

						// Set Stdout to variable.
						cmd.Stdout = &stdOut

						// Run execution of the current command.
						if err := cmd.Run(); err != nil {
							// If error was happened, add record with Stderr to fail results.
							r.Output = err.Error()
							// Add record with an error message to fail results.
							failResults = append(failResults, r)
						} else {
							// Check, if the current command want to print output.
							if t.IsPrintOutput {
								// Add record with Stdout to success results.
								r.Output = stdOut.String()
							}

							// Add record with command info to success results.
							successResults = append(successResults, r)
						}
					}(t)
				}
			}

			// Blocks all goroutines and wait the running commands.
			wg.Wait()

			// Check, if sequential queue have tasks.
			if len(app.Queues.Sequential) > 0 {
				// Loop for the all tasks in the sequential queue.
				for _, t := range app.Queues.Sequential {
					// Check, if current task have commands to execute.
					if t.Exec == nil || len(t.Exec) == 0 {
						return fmt.Errorf("can't run task without any commands to execute, see: %s", WikiPageURL)
					}

					// Add 1 to the wait group.
					wg.Add(1)

					// Create a new buffer for Stdout.
					var stdOut bytes.Buffer

					// Generate a new ID for current command.
					id, _ := gosl.RandomString(8)

					// Create a new struct for results.
					r := &Result{
						ID:          id,
						Name:        t.Name,
						Description: t.Description,
					}

					// Create a process for execute the current command.
					cmd := exec.Command(t.Exec[0], t.Exec[1:]...)

					// Check, if the current command want to run with sudo.
					if t.IsSudo {
						// Create a process for execute the current command as the sudo user.
						cmd = exec.Command("sudo", t.Exec...)
					}

					// Set Stdout to variable.
					cmd.Stdout = &stdOut

					// Run execution of the current command.
					if err := cmd.Run(); err != nil {
						// If error was happened, add record with Stderr to fail results.
						r.Output = err.Error()

						// Add record with an error message to fail results.
						failResults = append(failResults, r)
					} else {
						// Check, if the current command want to print output.
						if t.IsPrintOutput {
							// Add record with Stdout to success results.
							r.Output = stdOut.String()
						}

						// Add record with command info to success results.
						successResults = append(successResults, r)
					}

					// Don't forget to down the current wait group's count before goroutine is done.
					wg.Done()
				}
			}

			// Blocks all goroutines and wait the running commands.
			wg.Wait()

			// Check, if success results is not empty.
			if len(successResults) > 0 {
				printStyled(
					"Success commands:",
					"info",
					"margin-top-bottom",
				)

				// Loop for all success results.
				for _, result := range successResults {
					printStyled(
						fmt.Sprintf(
							"%s [name: %s, description: %s]",
							result.ID, result.Name, result.Description,
						),
						"success",
						"margin-left",
					)
					if result.Output != "" {
						printStyled(
							strings.TrimSuffix(result.Output, "\n"),
							"warning",
							"margin-left-2",
						)
					}
				}
			}

			// Check, if fail results is not empty.
			if len(failResults) > 0 {
				printStyled(
					"Fail commands:",
					"info",
					"margin-top-bottom",
				)

				// Loop for all fail results.
				for _, result := range failResults {
					printStyled(
						fmt.Sprintf(
							"%s [name: %s, description: %s]",
							result.ID, result.Name, result.Description,
						),
						"error",
						"margin-left",
					)
					printStyled(
						strings.TrimSuffix(result.Output, "\n"),
						"warning",
						"margin-left-2",
					)
				}
			}
		}
	}

	printStyled(
		fmt.Sprintf(
			"All done! 🎉 Time elapsed: %.2fs",
			time.Since(start).Seconds(),
		),
		"",
		"margin-top-bottom",
	)

	return nil
}
