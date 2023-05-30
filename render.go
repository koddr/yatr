package main

import "fmt"

// render provides the rendering process for results.
func (r *Results) render() {
	// Check, if success tasks queue have results.
	if len(r.Success) > 0 {
		// Print info message with count.
		printStyled(fmt.Sprintf("Success commands (%d):", len(r.Success)), successHeader)

		// Loop for all success results.
		for _, v := range r.Success {
			// Print info for a single success task.
			printStyled(fmt.Sprintf(
				"%s %s [name: %s, description: %s]",
				successMark, v.ID, applyStyle(v.Name, commandInfo), applyStyle(v.Description, commandInfo),
			), taskInfo)
		}
	}

	// Check, if fail tasks queue have results.
	if len(r.Fail) > 0 {
		// Print info message with count.
		printStyled(fmt.Sprintf("Fail commands (%d):", len(r.Fail)), failHeader)

		// Loop for all fail results.
		for _, v := range r.Fail {
			// Print info for a single fail task.
			printStyled(fmt.Sprintf(
				"%s %s [name: %s, description: %s]",
				failMark, v.ID, applyStyle(v.Name, commandInfo), applyStyle(v.Description, commandInfo),
			), taskInfo)

			// Print stderr output for a single fail task.
			printStyled(v.Output, outputMessage)
		}
	}
}
