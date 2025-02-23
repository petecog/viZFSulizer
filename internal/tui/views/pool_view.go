package views

import (
	"fmt"
	"strings"

	"github.com/petecog/vizfsulizer/internal/tui/styles"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

type PoolView struct {
	pools []*zfs.Pool
}

func NewPoolView() *PoolView {
	return &PoolView{}
}

func (pv *PoolView) Update(pools []*zfs.Pool) {
	pv.pools = pools
}

func (pv *PoolView) Render() string {
	if len(pv.pools) == 0 {
		return "No pools found"
	}

	var sb strings.Builder
	sb.WriteString(styles.Title.Render("ZFS Pools") + "\n\n")

	for _, pool := range pv.pools {
		poolContent := fmt.Sprintf("Pool: %s [%s]\n%s",
			styles.PoolName.Render(pool.Name),
			renderStatus(pool.Status),
			renderVDev(pool.RootVDev, 0))

		// Add Cache devices if present
		if pool.Cache != nil {
			poolContent += renderVDev(pool.Cache, 0)
		}

		// Add SLOG devices if present
		if pool.Slog != nil {
			poolContent += renderVDev(pool.Slog, 0)
		}

		boxedPool := styles.PoolBox.Render(poolContent)
		sb.WriteString(boxedPool + "\n\n")
	}

	// Add help text at the bottom
	sb.WriteString(styles.HelpText.Render("Press 'q' to quit, arrow keys to navigate"))
	return sb.String()
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
