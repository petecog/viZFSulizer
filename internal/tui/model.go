package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/petecog/vizfsulizer/internal/tui/views"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

// Model represents the main application state and handles the core UI logic.
// It manages the viewport, pool view, and pool selection state.
type Model struct {
	viewport viewport.Model  // Manages scrollable view area
	poolView *views.PoolView // Handles pool visualization
	pools    []*zfs.Pool     // List of ZFS pools to display
	selected int             // Currently selected pool index
}

// NewModel creates and initializes a new Model with default values.
// It sets up the viewport with zero initial size (will be updated later)
// and creates a new PoolView instance.
//
// Returns:
//   - Model: A new Model instance ready for use
func NewModel() Model {
	m := Model{
		viewport: viewport.New(0, 0), // Start with zero size, will be updated
		poolView: views.NewPoolView(),
		selected: 0,
	}
	return m
}

// Init implements tea.Model and returns the initial command to fetch pool data.
// This is called once when the program starts.
//
// Returns:
//   - tea.Cmd: Command to fetch initial pool data
func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		pools, _ := zfs.GetPools()
		return pools
	}
}

// Update implements tea.Model and handles all state updates.
// It processes different types of messages:
//   - WindowSizeMsg: Updates viewport dimensions
//   - KeyMsg: Handles keyboard input for navigation and quitting
//   - []*zfs.Pool: Updates pool data and view
//
// Parameters:
//   - msg: The message to process
//
// Returns:
//   - tea.Model: Updated model
//   - tea.Cmd: Any command to execute
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

// View implements tea.Model and returns the string to be displayed.
// It delegates to the viewport's View method to handle scrolling
// and content display.
//
// Returns:
//   - string: The complete rendered view
func (m Model) View() string {
	return m.viewport.View()
}
