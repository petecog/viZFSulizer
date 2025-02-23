// cmd/vizfsulizer/main.go
package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/peterjwest/vizfsulizer/internal/tui"
)

func main() {
    p := tea.NewProgram(tui.NewModel())
    if err := p.Start(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

// internal/tui/model.go
package tui

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Model struct {
    currentView string
}

func NewModel() Model {
    return Model{
        currentView: "Welcome to viZFSulizer!",
    }
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m Model) View() string {
    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("205")).
        Bold(true).
        Padding(1, 2)

    return style.Render(m.currentView) + "\n\nPress 'q' to quit\n"
}