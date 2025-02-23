package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	StatusOnline = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")). // Green
			Bold(true)

	StatusDegraded = lipgloss.NewStyle().
			Foreground(lipgloss.Color("3")). // Yellow
			Bold(true)

	StatusFaulted = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")). // Red
			Bold(true)

	PoolName = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4")). // Blue
			Bold(true)

	VDevType = lipgloss.NewStyle().
			Foreground(lipgloss.Color("5")) // Magenta

	TreeBranch = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // Gray
)
