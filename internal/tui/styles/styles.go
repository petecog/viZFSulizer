package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	StatusOnline = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorGreen)).
			Bold(true)

	StatusDegraded = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorYellow)).
			Bold(true)

	StatusFaulted = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorRed)).
			Bold(true)

	PoolName = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorBlue)).
			Bold(true)

	VDevType = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorDarkGold))

	TreeBranch = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorGray))

	// VDev related styles
	VDevName = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorSkyBlue))

	// Optional related styles for consistency
	VDevSize = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPaleGreen))

	VDevPath = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorDisk))
)
