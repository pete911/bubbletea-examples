package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Title struct {
	style      lipgloss.Style
	itemStyle  lipgloss.Style
	keyStyle   lipgloss.Style
	valueStyle lipgloss.Style
	items      []TitleItem
}

type TitleItem struct {
	Key   string
	Value string
}

func NewTitle(items []TitleItem) Title {
	return Title{
		style:      lipgloss.NewStyle().Align(lipgloss.Left).Border(lipgloss.NormalBorder(), true, false, true, false),
		itemStyle:  lipgloss.NewStyle().Padding(0, 5, 0, 0).Align(lipgloss.Left),
		keyStyle:   KeyFont.Copy().Padding(0, 1, 0, 0),
		valueStyle: ValueFont.Copy().Padding(),
		items:      items,
	}
}

func (t Title) View(width int) string {
	width = width - t.style.GetHorizontalFrameSize()
	var renderedItems []string
	for _, v := range t.items {
		keyValue := lipgloss.JoinHorizontal(lipgloss.Top, t.keyStyle.Render(v.Key), t.valueStyle.Render(v.Value))
		renderedItems = append(renderedItems, t.itemStyle.Render(keyValue))
	}

	bar := lipgloss.JoinHorizontal(lipgloss.Top, renderedItems...)
	return t.style.Width(width).Render(bar)
}
