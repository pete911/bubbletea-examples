package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ping          Ping
	Title         Title
	Content       string
	Footer        Footer
	style         lipgloss.Style
	width, height int
}

func NewModel() Model {
	return Model{
		ping:    NewPing(),
		Title:   NewTitle([]TitleItem{{Key: "Name", Value: "BubbleTea example"}, {Key: "Version", Value: "0.0.1"}}),
		Content: "...",
		Footer:  NewFooter("Demo BubbleTea, press Ctrl+c to exit"),
		style:   lipgloss.NewStyle().Padding(0, 1),
	}
}

func (m Model) Init() tea.Cmd {
	return m.ping.Cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case Response:
		m.Content = string(msg)
		return m, m.ping.Cmd

	case ErrorResponse:
		m.Content = string(msg)
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}

	return m, nil
}

func (m Model) View() string {
	width := m.width - m.style.GetHorizontalFrameSize()
	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, m.Title.View(width), m.Content, m.Footer.View(width)))
}
