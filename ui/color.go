package ui

import "github.com/charmbracelet/lipgloss"

var (
	KeyFont   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#484848", Dark: "#e1e1e1"})
	ValueFont = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#7d7d7d", Dark: "#b4b4b4"})
)
