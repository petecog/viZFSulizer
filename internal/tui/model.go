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
	selected int // Currently selected pool index
}

func NewModel() Model {
	m := Model{
		viewport: viewport.New(0, 0), // Start with zero size, will be updated
		poolView: views.NewPoolView(),
		selected: 0,
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
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "right", "l":
			if len(m.pools) > 0 {
				m.selected = (m.selected + 1) % len(m.pools)
				m.poolView.SetSelected(m.selected)
				m.viewport.SetContent(m.poolView.Render())
			}
		case "shift+tab", "left", "h":
			if len(m.pools) > 0 {
				m.selected = (m.selected - 1 + len(m.pools)) % len(m.pools)
				m.poolView.SetSelected(m.selected)
				m.viewport.SetContent(m.poolView.Render())
			}
		}

	case []*zfs.Pool:
		m.pools = msg
		m.poolView.Update(msg)
		m.poolView.SetSelected(m.selected)
		m.viewport.SetContent(m.poolView.Render())
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.viewport.View()
}
