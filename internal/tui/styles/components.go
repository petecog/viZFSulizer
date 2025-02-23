package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (

	// Component styles
	PoolBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color(ColorPrimary)).
		Padding(1).
		Width(76)

	// VDevBox defines box styling for VDev related content.
	// It applies a custom border with VDev-specific color,
	// adds horizontal internal padding of 1 unit,
	// and sets a left margin of 2 units.
	VDevBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color(ColorVDev)).
		Padding(0, 1).
		MarginLeft(2)

	DiskInfo = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorDisk))

	ResilvingProgress = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorWarning)).
				Bold(true)

	Capacity = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorEmphasis))
)
