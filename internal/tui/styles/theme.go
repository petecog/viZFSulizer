package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Box styling
	BoxBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	// Component styles
	PoolBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color("4")).
		Padding(1).
		Width(76)

	VDevBox = lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color("5")).
		MarginLeft(2).
		Padding(0, 1)

	// Selection highlighting
	Selected = lipgloss.NewStyle().
			Background(lipgloss.Color("4")).
			Foreground(lipgloss.Color("0"))

	// Headers and titles
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6")).
		MarginLeft(2)

	// Help text
	HelpText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			AlignHorizontal(lipgloss.Right)

	// Tab styling
	TabActive = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("4")).
			Background(lipgloss.Color("0"))

	TabInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))
)
