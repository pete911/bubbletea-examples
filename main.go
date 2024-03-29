package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pete911/bubbletea-examples/ui"
	"os"
)

func main() {
	if _, err := tea.NewProgram(ui.NewModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
