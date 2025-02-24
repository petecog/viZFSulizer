package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Start initializes and runs the Terminal User Interface (TUI) program.
// It sets up the main application model and configures the Bubble Tea
// program with appropriate options for terminal handling.
//
// The function:
// 1. Creates a new application model
// 2. Initializes the Bubble Tea program with:
//   - Alternate screen buffer for clean UI
//   - Mouse support for enhanced interaction
//
// 3. Runs the main program loop
//
// Returns:
//   - error: Any error that occurred during program execution
//
// Example usage:
//
//	if err := tui.Start(); err != nil {
//	    log.Fatal("Failed to start TUI:", err)
//	}
func Start() error {
	// Initialize model with default state
	model := NewModel()

	// Create program with options for better UI experience
	fmt.Println("viZFSulizer Starting...")
	p := tea.NewProgram(model,
		tea.WithAltScreen(),       // Use alternate screen for clean UI
		tea.WithMouseCellMotion(), // Enable mouse interaction
	)

	// Run the program and return any errors
	_, err := p.Run()
	return err
}
