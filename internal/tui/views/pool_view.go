package views

import (
	"fmt"
	"strings"

	"github.com/petecog/vizfsulizer/internal/tui/styles"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

type PoolView struct {
	pools        []*zfs.Pool
	selected     int
	vdevExpanded bool
	expandedVDev string
	vdevView     *VDevView
}

func NewPoolView() *PoolView {
	return &PoolView{
		selected: 0,
	}
}

func (pv *PoolView) Update(pools []*zfs.Pool) {
	pv.pools = pools
	if len(pools) > 0 {
		// Create/update VDevView with currently selected pool
		pv.vdevView = NewVDevView(pools[pv.selected])
	}
}

func (pv *PoolView) SetSelected(idx int) {
	pv.selected = idx
	if len(pv.pools) > 0 {
		// Update VDevView when selection changes
		pv.vdevView = NewVDevView(pv.pools[idx])
	}
}

func (pv *PoolView) ToggleVDevExpanded() {
	pv.vdevExpanded = !pv.vdevExpanded
}

func (pv *PoolView) Render() string {
	if len(pv.pools) == 0 {
		return "No pools found"
	}

	var sb strings.Builder

	// Render tabs
	sb.WriteString(renderTabs(pv.pools, pv.selected) + "\n\n")

	// Render selected pool overview
	pool := pv.pools[pv.selected]
	worstStatus := pool.GetWorstStatus()
	poolOverview := fmt.Sprintf("Pool: %s [%s]",
		styles.PoolName.Render(pool.Name),
		renderStatus(worstStatus))

	sb.WriteString(styles.GetStatusBorderStyle(worstStatus).Render(poolOverview) + "\n\n")

	// Modify VDev rendering based on expanded state
	var vdevContent string
	if pv.vdevExpanded {
		vdevContent = pv.vdevView.RenderDetailed(pv.expandedVDev)
	} else {
		vdevContent = pv.vdevView.RenderCompact()
	}
	sb.WriteString(vdevContent)

	// Update help text to show VDEV navigation
	helpText := "Tab/Arrow Keys to switch pools • "
	helpText += "↑/↓ to navigate VDEVs • "
	helpText += "Enter to expand/collapse • "
	helpText += "q to quit"
	sb.WriteString("\n" + styles.HelpText.Render(helpText))

	return sb.String()
}

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

// renderStatus applies styling to a ZFS VDev status string based on its state.
// It returns a styled string representation of the VDev status.
//
// The following styles are applied:
//   - Online: Uses StatusOnline style
//   - Degraded: Uses StatusDegraded style
//   - Faulted: Uses StatusFaulted style
//   - Other statuses: Returns unmodified status string
//
// Parameters:
//   - status: zfs.VDevStatus representing the current state of a VDev
//
// Returns:
//   - string: A styled string representation of the VDev status
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

// Add these methods after the existing PoolView methods

// NavigateVDevs moves the VDEV selection up or down
// direction: -1 for up, 1 for down
func (pv *PoolView) NavigateVDevs(direction int) {
	// Get flat list of navigable VDEVs
	vdevs := pv.getNavigableVDevs()
	if len(vdevs) == 0 {
		return
	}

	// Find current index
	currentIdx := -1
	for i, v := range vdevs {
		if v.Name == pv.expandedVDev {
			currentIdx = i
			break
		}
	}

	// Calculate new index with wrapping
	if currentIdx == -1 {
		currentIdx = 0
	} else {
		currentIdx = (currentIdx + direction + len(vdevs)) % len(vdevs) // Fixed missing parenthesis
	}

	// Update selected VDEV
	pv.expandedVDev = vdevs[currentIdx].Name
}

// getNavigableVDevs returns a flat list of all VDEVs in the current pool
func (pv *PoolView) getNavigableVDevs() []*zfs.VDev {
	var vdevs []*zfs.VDev
	if pv.pools == nil || len(pv.pools) == 0 {
		return vdevs
	}

	pool := pv.pools[pv.selected]
	// Collect all VDEVs in a flat list for navigation
	if pool.RootVDev != nil {
		vdevs = append(vdevs, pv.collectVDevs(pool.RootVDev)...)
	}
	if pool.Cache != nil {
		vdevs = append(vdevs, pv.collectVDevs(pool.Cache)...)
	}
	if pool.Slog != nil {
		vdevs = append(vdevs, pv.collectVDevs(pool.Slog)...)
	}
	return vdevs
}

// collectVDevs helper function to recursively collect VDEVs
func (pv *PoolView) collectVDevs(vdev *zfs.VDev) []*zfs.VDev {
	var vdevs []*zfs.VDev
	vdevs = append(vdevs, vdev)
	for _, child := range vdev.Children {
		vdevs = append(vdevs, pv.collectVDevs(child)...)
	}
	return vdevs
}
