package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/knadh/koanf/v2"
)

func main() {
	// Set flags.
	path := flag.String("p", "./tasks.json", "set a path of the tasks file")

	// Parse all flags.
	flag.Parse()

	// Create the koanf instance for a root path.
	k := koanf.New(".")

	// Parse all tasks from the given file.
	tt, err := newParser(path, k)
	if err != nil {
		printStyled(fmt.Sprintf("%s %v", failMark, err.Error()), failHeader)
		os.Exit(0)
	}

	// Create a new queue for async and sequential tasks.
	q, err := newQueue(tt)
	if err != nil {
		printStyled(fmt.Sprintf("%s %v", failMark, err.Error()), failHeader)
		os.Exit(0)
	}

	// Print the welcome message.
	printStyled("🏃 Welcome to yatr (Yet Another Task Runner)!", welcomeHeader)

	// Print tasks info message.
	printStyled(fmt.Sprintf(
		"Running tasks set from %s [name: %s, description: %s]\n%d tasks in queue (async: %d, sequential: %d)... please wait!",
		*path, tt.Name, tt.Description, len(tt.Tasks), len(q.AsyncQueue), len(q.SequentialQueue),
	), warningHeader)

	// Start timer.
	start := time.Now()

	// Run all tasks and print results.
	q.runTasks().render()

	// Print the executing message.
	printStyled(fmt.Sprintf(
		"🎉 Done! Time spent executing %d tasks: %vs",
		len(tt.Tasks), time.Since(start).Seconds(),
	), doneHeader)
}
