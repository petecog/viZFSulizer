package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/petecog/vizfsulizer/internal/zfs" // Fix import path
)

// Theme definitions for the TUI components.
// This file contains all the styling and theme related constants and functions
// using the Lipgloss styling library for terminal UI rendering.
var (
	// BoxBorder defines the default border characters used for boxes and panels.
	// Uses Unicode box-drawing characters to create clean, professional borders.
	BoxBorder = lipgloss.Border{
		Top:         "─", // Horizontal line for top border
		Bottom:      "─", // Horizontal line for bottom border
		Left:        "│", // Vertical line for left border
		Right:       "│", // Vertical line for right border
		TopLeft:     "╭", // Corner piece for top-left
		TopRight:    "╮", // Corner piece for top-right
		BottomLeft:  "╰", // Corner piece for bottom-left
		BottomRight: "╯", // Corner piece for bottom-right
	}

	// PoolBox defines the style for the main pool container.
	// Uses blue borders (ANSI color 4) with padding and fixed width
	// to create a consistent layout for pool information.
	PoolBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color("4")). // Blue border
		Padding(1).
		Width(76)

	// VDevBox defines the style for virtual device containers.
	// Uses purple borders (ANSI color 5) with left margin for hierarchy
	// and minimal padding for compact display.
	VDevBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color("5")). // Purple border
		MarginLeft(2).
		Padding(0, 1)

	// Selected defines the highlight style for selected items.
	// Uses blue background (ANSI color 4) with black text (ANSI color 0)
	// to create high contrast for selected elements.
	Selected = lipgloss.NewStyle().
			Background(lipgloss.Color("4")). // Blue background
			Foreground(lipgloss.Color("0"))  // Black text

	// Title defines the style for section headers and titles.
	// Uses cyan text (ANSI color 6) in bold with left margin
	// for visual hierarchy.
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6")). // Cyan text
		MarginLeft(2)

	// HelpText defines the style for user instructions and help messages.
	// Uses gray text (ANSI color 8) aligned to the right
	// for subtle but accessible help information.
	HelpText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")). // Gray text
			AlignHorizontal(lipgloss.Right)

	// TabActive defines the style for the currently selected tab.
	// Uses blue text (ANSI color 4) on black background (ANSI color 0)
	// with bold for emphasis.
	TabActive = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("4")). // Blue text
			Background(lipgloss.Color("0"))  // Black background

	// TabInactive defines the style for non-selected tabs.
	// Uses gray text (ANSI color 8) for de-emphasized display.
	TabInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // Gray text
)

// GetStatusBorderStyle returns a border style based on the VDev status.
// The border color indicates the health state:
//   - Green (ANSI color 2) for ONLINE status
//   - Yellow (ANSI color 3) for DEGRADED status
//   - Red (ANSI color 1) for FAULTED status
//
// Parameters:
//   - status: The VDev status to determine the border color
//
// Returns:
//   - lipgloss.Style: A styled border matching the status severity
func GetStatusBorderStyle(status zfs.VDevStatus) lipgloss.Style {
	var color string
	switch status {
	case zfs.VDevStatusFaulted:
		color = "1" // Red for critical failure
	case zfs.VDevStatusDegraded:
		color = "3" // Yellow for warning
	default:
		color = "2" // Green for healthy
	}

	return lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color(color)).
		Padding(0, 1)
}
