package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Start initializes the TUI program and runs it
func Start() error {
	// Initialize model
	model := NewModel()

	// Create program
	fmt.Println("viZFSulizer Starting...")
	p := tea.NewProgram(model,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	// Run program
	_, err := p.Run()
	return err
}
