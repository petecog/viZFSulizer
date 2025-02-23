package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Header and title styling
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(ColorEmphasis)).
		MarginLeft(2)

	// Help text
	HelpText = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			AlignHorizontal(lipgloss.Right)

	// Tab styling
	TabActive = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorPrimary)).
			Background(lipgloss.Color(ColorInverse))

	TabInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted))

	// Selection highlighting
	SelectedText = lipgloss.NewStyle().
			Background(lipgloss.Color("4")).
			Foreground(lipgloss.Color("0"))
)
