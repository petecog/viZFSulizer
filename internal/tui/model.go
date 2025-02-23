package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/petecog/vizfsulizer/internal/tui/views"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

type Model struct {
	viewport viewport.Model
	poolView *views.PoolView
	pools    []*zfs.Pool
}

func NewModel() Model {
	m := Model{
		viewport: viewport.New(0, 0), // Start with zero size, will be updated
		poolView: views.NewPoolView(),
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		pools, _ := zfs.GetPools()
		return pools
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case []*zfs.Pool:
		m.pools = msg
		m.poolView.Update(msg)
		m.viewport.SetContent(m.poolView.Render())
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.viewport.View()
}
