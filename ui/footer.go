package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Footer struct {
	style   lipgloss.Style
	content string
}

func NewFooter(content string) Footer {
	return Footer{
		style:   ValueFont.Copy().Align(lipgloss.Left).Border(lipgloss.NormalBorder(), true, false, true, false),
		content: content,
	}
}

func (f Footer) View(width int) string {
	return f.style.Width(width).Render(f.content)
}
