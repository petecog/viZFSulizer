package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

// Box styling
var BoxBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "╰",
	BottomRight: "╯",
}

func GetStatusBorderStyle(status zfs.VDevStatus) lipgloss.Style {
	var color string
	switch status {
	case zfs.VDevStatusFaulted:
		color = ColorError
	case zfs.VDevStatusDegraded:
		color = ColorWarning
	default:
		color = ColorSuccess
	}

	return lipgloss.NewStyle().
		Border(BoxBorder).
		BorderForeground(lipgloss.Color(color)).
		Padding(0, 1)
}
