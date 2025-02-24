package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Style definitions for the TUI components.
// Each style defines colors and text formatting for different UI elements
// using the Lipgloss styling library.
var (
	// StatusOnline defines the style for healthy/online components
	// Uses bright green (ANSI color 2) to indicate normal operation
	StatusOnline = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")). // Green - indicates healthy state
			Bold(true)

	// StatusDegraded defines the style for components with reduced functionality
	// Uses yellow (ANSI color 3) to indicate warning state
	StatusDegraded = lipgloss.NewStyle().
			Foreground(lipgloss.Color("3")). // Yellow - indicates warning/degraded
			Bold(true)

	// StatusFaulted defines the style for failed components
	// Uses red (ANSI color 1) to indicate critical failure
	StatusFaulted = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")). // Red - indicates failure/error
			Bold(true)

	// PoolName defines the style for ZFS pool names
	// Uses blue (ANSI color 4) to make pool names stand out
	PoolName = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4")). // Blue - highlights pool names
			Bold(true)

	// VDevType defines the style for virtual device type labels
	// Uses magenta (ANSI color 5) to distinguish device types
	VDevType = lipgloss.NewStyle().
			Foreground(lipgloss.Color("5")) // Magenta - shows device types

	// TreeBranch defines the style for the tree view connection lines
	// Uses gray (ANSI color 8) to create subtle connection lines
	TreeBranch = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // Gray - subtle tree structure
)
