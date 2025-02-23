package views

import (
	"fmt"
	"strings"

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
	for _, pool := range pv.pools {
		sb.WriteString(fmt.Sprintf("Pool: %s [%s]\n", pool.Name, pool.Status))
		sb.WriteString(renderVDev(pool.RootVDev, 0))
	}
	return sb.String()
}

func renderVDev(vdev *zfs.VDev, depth int) string {
	indent := strings.Repeat("  ", depth)
	result := fmt.Sprintf("%s├─ %s (%s) [%s]\n", indent, vdev.Name, vdev.Type, vdev.Status)

	for _, child := range vdev.Children {
		result += renderVDev(child, depth+1)
	}
	return result
}
