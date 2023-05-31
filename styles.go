package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Set colors.
	successColor = lipgloss.AdaptiveColor{Light: "#16a34a", Dark: "#4ade80"}
	warningColor = lipgloss.AdaptiveColor{Light: "#ca8a04", Dark: "#facc15"}
	failColor    = lipgloss.AdaptiveColor{Light: "#dc2626", Dark: "#f87171"}
	outputColor  = lipgloss.AdaptiveColor{Light: "#4b5563", Dark: "#9ca3af"}

	// Set headers styles.
	welcomeHeader = lipgloss.NewStyle().Foreground(outputColor).Margin(1, 0, 0).Bold(true)
	successHeader = lipgloss.NewStyle().Foreground(successColor).Margin(0, 0, 1)
	warningHeader = lipgloss.NewStyle().Foreground(warningColor).Margin(1, 0)
	failHeader    = lipgloss.NewStyle().Foreground(failColor).Margin(1, 0)
	doneHeader    = lipgloss.NewStyle().Foreground(outputColor).Margin(0, 0, 1).Bold(true)

	// Set task info style.
	taskInfo = lipgloss.NewStyle().Margin(0)

	// Set command info style.
	commandInfo = lipgloss.NewStyle().Foreground(outputColor)

	// Set output style.
	outputMessage = lipgloss.NewStyle().Foreground(outputColor).Margin(1, 0)

	// Create elements with styles.
	successMark = lipgloss.NewStyle().Foreground(successColor).SetString("✓")
	failMark    = lipgloss.NewStyle().Foreground(failColor).SetString("✕")
)

// applyStyle applies style for the given string by a lipgloss.Style template.
func applyStyle(s string, ls lipgloss.Style) string {
	return ls.Render(s)
}

// printStyled prints styled output for the given string by a lipgloss.Style template.
func printStyled(s string, ls lipgloss.Style) {
	fmt.Println(ls.Render(s))
}
