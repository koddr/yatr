package main

import (
	_ "embed"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/koddr/gosl"
)

// WikiPageURL set a full path to the Wiki page.
const WikiPageURL string = "https://github.com/koddr/yatr/wiki"

var (
	// Flags for init example tasks file.
	initFlag bool

	// Flag for set up a path to the tasks file.
	pathFlag string

	//go:embed examples/tasks.yaml
	embedYAMLTasksFile []byte
)

// printStyled prints styled output for the given string by a lipgloss.Style template.
func printStyled(s, state, style string) {
	// Set lipgloss colors.
	successColor := lipgloss.AdaptiveColor{Light: "#16a34a", Dark: "#4ade80"}
	warningColor := lipgloss.AdaptiveColor{Light: "#ca8a04", Dark: "#facc15"}
	errorColor := lipgloss.AdaptiveColor{Light: "#dc2626", Dark: "#f87171"}
	infoColor := lipgloss.AdaptiveColor{Light: "#4b5563", Dark: "#9ca3af"}

	// Create a new lipgloss style instance.
	lg := lipgloss.NewStyle()

	// Switch between states.
	switch state {
	case "info":
		state = lg.Foreground(infoColor).SetString("– ").String()
	case "success":
		state = lg.Foreground(successColor).SetString("✓ ").String()
	case "error":
		state = lg.Foreground(errorColor).SetString("✕ ").String()
	case "warning":
		state = lg.Foreground(warningColor).SetString("‼ ").String()
	}

	// Switch between styles.
	switch style {
	case "margin-top-bottom":
		s = gosl.RenderStyled(gosl.Concat(state, s), lg.MarginTop(1).MarginBottom(1))
	case "margin-top":
		s = gosl.RenderStyled(gosl.Concat(state, s), lg.MarginTop(1))
	case "margin-bottom":
		s = gosl.RenderStyled(gosl.Concat(state, s), lg.MarginBottom(1))
	case "margin-left":
		s = gosl.RenderStyled(gosl.Concat(state, s), lg.MarginLeft(1))
	case "margin-left-2":
		s = gosl.RenderStyled(gosl.Concat(state, s), lg.MarginLeft(2))
	default:
		s = gosl.Concat(state, s)
	}

	// Print styled output.
	fmt.Println(s)
}
