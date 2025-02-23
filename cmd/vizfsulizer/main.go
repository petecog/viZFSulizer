// cmd/vizfsulizer/main.go
package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/petecog/vizfsulizer/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.NewModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
