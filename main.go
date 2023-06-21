package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Set flags.
	flag.BoolVar(&initFlag, "i", false, "generate an example 'tasks.yml' file in the given format")
	flag.StringVar(&pathFlag, "p", "", "set a path to the tasks file")

	// Parse all flags.
	flag.Parse()

	if initFlag {
		// Create the tasks file with the given format in the current dir.
		if err := os.WriteFile(filepath.Clean("./tasks.yml"), embedYAMLTasksFile, 0o600); err != nil {
			printStyled(
				fmt.Sprintf("There was an error with generating example tasks file: %v", err),
				"error",
				"",
			)
			os.Exit(0)
		}
	} else {
		// App initialization.
		app, err := initialize()
		if err != nil {
			printStyled(
				fmt.Sprintf("There was an error with initialize app: %v", err),
				"error",
				"",
			)
			os.Exit(0)
		}

		// App start.
		if err = app.start(); err != nil {
			printStyled(
				fmt.Sprintf("There was an error with starting app: %v", err),
				"error",
				"",
			)
			os.Exit(0)
		}
	}
}
