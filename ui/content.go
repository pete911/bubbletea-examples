package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Content struct {
	style   lipgloss.Style
	content string
}

func NewContent(content string) Footer {
	return Footer{
		style:   ValueFont.Copy().Align(lipgloss.Left),
		content: content,
	}
}

func (c Content) View(width int) string {
	return c.style.Width(width).Render(c.content)
}
