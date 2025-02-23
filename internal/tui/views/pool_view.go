package views

import (
	"fmt"
	"strings"

	"github.com/petecog/vizfsulizer/internal/tui/styles"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

type PoolView struct {
	pools    []*zfs.Pool
	selected int
}

func NewPoolView() *PoolView {
	return &PoolView{}
}

func (pv *PoolView) Update(pools []*zfs.Pool) {
	pv.pools = pools
}

func (pv *PoolView) SetSelected(idx int) {
	pv.selected = idx
}

func (pv *PoolView) Render() string {
	if len(pv.pools) == 0 {
		return "No pools found"
	}

	var sb strings.Builder

	// Render tabs
	sb.WriteString(renderTabs(pv.pools, pv.selected) + "\n\n")

	// Render selected pool
	pool := pv.pools[pv.selected]
	poolContent := fmt.Sprintf("Pool: %s [%s]\n%s",
		styles.PoolName.Render(pool.Name),
		renderStatus(pool.Status),
		renderVDev(pool.RootVDev, 0))

	if pool.Cache != nil {
		poolContent += renderVDev(pool.Cache, 0)
	}
	if pool.Slog != nil {
		poolContent += renderVDev(pool.Slog, 0)
	}

	boxedPool := styles.PoolBox.Render(poolContent)
	sb.WriteString(boxedPool + "\n\n")

	// Update help text to include tab navigation
	sb.WriteString(styles.HelpText.Render("Tab/Arrow Keys to switch pools • q to quit"))
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

func renderVDev(vdev *zfs.VDev, depth int) string {
	content := fmt.Sprintf("%s %s %s [%s]",
		styles.TreeBranch.Render(strings.Repeat("  ", depth)+"├─"),
		vdev.Name,
		styles.VDevType.Render("("+vdev.Type+")"),
		renderStatus(vdev.Status))

	if len(vdev.Children) > 0 {
		childContent := ""
		for _, child := range vdev.Children {
			childContent += renderVDev(child, depth+1)
		}
		content = styles.VDevBox.Render(content + "\n" + childContent)
	}

	return content + "\n"
}

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
