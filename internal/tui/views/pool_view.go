package views

import (
	"fmt"
	"strings"

	"github.com/petecog/vizfsulizer/internal/tui/styles"
	"github.com/petecog/vizfsulizer/internal/zfs"
	"github.com/petecog/vizfsulizer/internal/zfs/status"
)

// PoolView represents the visual component for displaying ZFS pool information.
// It maintains the current state of pools and handles their rendering.
// The view supports multiple pools with tab-based navigation and detailed
// status information for each pool's virtual devices (VDevs).
// TODO? Split out maintain current state and rendering into separate structs?
type PoolView struct {
	pools    []*zfs.Pool      // List of ZFS pools to display
	selected int              // Index of currently selected pool
	analyzer *status.Analyzer // Tool for analyzing pool and VDev health
}

// NewPoolView creates and initializes a new PoolView with default values.
// It sets up a status analyzer for determining pool and VDev health states.
// The analyzer is used to recursively check the status of all devices
// in the pool hierarchy.
//
// Returns:
//   - *PoolView: A new PoolView instance ready for use
func NewPoolView() *PoolView {
	return &PoolView{
		analyzer: &status.Analyzer{},
	}
}

// Update refreshes the pool data stored in the PoolView.
// It takes a slice of Pool pointers and updates the internal state.
// This method should be called whenever the underlying pool data changes.
//
// Parameters:
//   - pools: New slice of Pool pointers to display
func (pv *PoolView) Update(pools []*zfs.Pool) {
	pv.pools = pools
}

// SetSelected updates the currently selected pool index.
// This determines which pool's details are displayed in the view.
// The index should be within the valid range of the pools slice.
//
// Parameters:
//   - idx: Index of the pool to select
func (pv *PoolView) SetSelected(idx int) {
	pv.selected = idx
}

// Render generates the complete string representation of the PoolView.
// It creates a formatted display including:
//   - A tab bar showing all available pools
//   - Detailed information about the selected pool
//   - Status indicators for all VDevs in the pool
//   - Help text showing available commands
//
// The output uses border styles based on pool health status and
// includes proper spacing and alignment for readability.
//
// Returns:
//   - string: The complete rendered view ready for display
//
// Example Output:
//
//	[ pool1 ]  pool2   pool3
//
//	Pool: pool1 [ONLINE]
//	├─ mirror-0 (mirror) [ONLINE]
//	│  ├─ sda (disk) [ONLINE]
//	│  └─ sdb (disk) [ONLINE]
//
//	Tab/Arrow Keys to switch pools • q to quit
func (pv *PoolView) Render() string {
	if len(pv.pools) == 0 {
		return "No pools found"
	}

	var sb strings.Builder

	// Render tabs
	sb.WriteString(renderTabs(pv.pools, pv.selected) + "\n\n")

	// Render selected pool
	pool := pv.pools[pv.selected]
	worstStatus := pv.analyzer.GetPoolWorstStatus(pool) // Use analyzer's GetPoolWorstStatus
	poolContent := fmt.Sprintf("Pool: %s [%s]\n%s",
		styles.PoolName.Render(pool.Name),
		renderStatus(worstStatus),
		renderVDev(pool.RootVDev, 0, pv.analyzer))

	if pool.Cache != nil {
		poolContent += renderVDev(pool.Cache, 0, pv.analyzer)
	}
	if pool.Slog != nil {
		poolContent += renderVDev(pool.Slog, 0, pv.analyzer)
	}

	boxedPool := styles.GetStatusBorderStyle(worstStatus).Render(poolContent)
	sb.WriteString(boxedPool + "\n\n")

	// Update help text to include tab navigation
	sb.WriteString(styles.HelpText.Render("Tab/Arrow Keys to switch pools • q to quit"))
	return sb.String()
}

// renderTabs creates a horizontal tab bar showing all pool names.
// It displays each pool name with appropriate styling to indicate
// the currently selected pool. This provides visual feedback for
// pool navigation.
//
// Parameters:
//   - pools: Slice of Pool pointers to render as tabs
//   - selected: Index of the currently selected pool
//
// Returns:
//   - string: A formatted string containing the tab bar
//
// Example:
//
//	[ pool1 ]  pool2   pool3
func renderTabs(pools []*zfs.Pool, selected int) string {
	var tabs []string
	for i, pool := range pools {
		tab := pool.Name
		if i == selected {
			tab = styles.TabActive.Render("[ " + tab + " ]")
		} else {
			tab = styles.TabInactive.Render("  " + tab + "  ")
		}
		tabs = append(tabs, tab)
	}
	return strings.Join(tabs, " ")
}

// renderVDev creates a string representation of a VDev and its children.
// It recursively renders the entire VDev tree with proper indentation and styling.
// Parameters:
//   - vdev: pointer to the VDev to render
//   - depth: current depth in the VDev tree for indentation
//   - analyzer: status analyzer for determining VDev health
//
// Returns a string containing the rendered VDev tree.
func renderVDev(vdev *zfs.VDev, depth int, analyzer *status.Analyzer) string {
	// Use the analyzer to get the worst status of the VDev
	worstStatus := analyzer.GetVDevWorstStatus(vdev)
	content := fmt.Sprintf("%s %s %s [%s]",
		styles.TreeBranch.Render(strings.Repeat("  ", depth)+"├─"),
		vdev.Name,
		styles.VDevType.Render("("+vdev.Type+")"),
		renderStatus(worstStatus))

	if len(vdev.Children) > 0 {
		childContent := ""
		for _, child := range vdev.Children {
			childContent += renderVDev(child, depth+1, analyzer)
		}
		content = styles.GetStatusBorderStyle(worstStatus).Render(content + "\n" + childContent)
	}

	return content + "\n"
}

// renderStatus converts a VDevStatus to a styled string representation.
// It applies different colors and styles based on the status value.
// Parameters:
//   - status: the VDevStatus to render
//
// Returns a styled string representing the status.
func renderStatus(status zfs.VDevStatus) string {
	switch status {
	case zfs.VDevStatusOnline:
		return styles.StatusOnline.Render(string(status))
	case zfs.VDevStatusDegraded:
		return styles.StatusDegraded.Render(string(status))
	case zfs.VDevStatusFaulted:
		return styles.StatusFaulted.Render(string(status))
	default:
		return string(status)
	}
}
