package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/petecog/vizfsulizer/internal/tui/styles"
	"github.com/petecog/vizfsulizer/internal/zfs"
)

// VDevView provides a visual representation of ZFS virtual devices
// and their current status in the terminal UI
type VDevView struct {
	pool  *zfs.Pool
	style lipgloss.Style
}

// NewVDevView creates a new VDev view instance with the specified pool
// and initializes the default styling
func NewVDevView(pool *zfs.Pool) *VDevView {
	return &VDevView{
		pool: pool,
	}
}

// RenderVDev creates a formatted view of a VDev and its children
func (v *VDevView) RenderVDev(vdev *zfs.VDev, depth int) string {
	worstStatus := v.pool.GetWorstStatus()
	var content strings.Builder

	// Basic VDev info with tree structure
	content.WriteString(fmt.Sprintf("%s %s %s [%s]",
		styles.TreeBranch.Render(strings.Repeat("  ", depth)+"├─"),
		vdev.Name,
		styles.VDevType.Render("("+string(vdev.Type)+")"),
		renderStatus(worstStatus)))

	// Add capacity info if available
	if vdev.TotalCapacity > 0 {
		content.WriteString(fmt.Sprintf("\nCapacity: %d/%d GB",
			vdev.UsedCapacity/1024/1024/1024,
			vdev.TotalCapacity/1024/1024/1024))
	}

	// Render disk information if any
	if len(vdev.Disks) > 0 {
		content.WriteString("\n" + v.renderDisks(vdev.Disks))
	}

	// Get children content before applying border
	var childrenContent string
	if len(vdev.Children) > 0 {
		var childBuilder strings.Builder
		for _, child := range vdev.Children {
			// Each child gets its own border
			childBox := v.RenderVDev(child, depth+1)
			childBuilder.WriteString("\n" + childBox)
		}
		childrenContent = childBuilder.String()
	}

	// Apply border to current VDev content
	boxed := styles.GetStatusBorderStyle(worstStatus).Render(content.String())

	// Add children after current VDev's border
	if childrenContent != "" {
		boxed += childrenContent
	}

	return boxed
}

// Render generates the complete VDev hierarchy view
func (v *VDevView) Render() string {
	var output string

	// Start with the root VDev
	if v.pool.RootVDev != nil {
		output += v.RenderVDev(v.pool.RootVDev, 0)
	}

	// Add Cache devices if present
	if v.pool.Cache != nil {
		output += v.RenderVDev(v.pool.Cache, 0)
	}

	// Add SLOG devices if present
	if v.pool.Slog != nil {
		output += v.RenderVDev(v.pool.Slog, 0)
	}

	return output
}

// RenderCompact shows a condensed view of all VDEVs
func (v *VDevView) RenderCompact() string {
	var output strings.Builder

	if v.pool.RootVDev != nil {
		output.WriteString(v.renderVDevCompact(v.pool.RootVDev))
	}

	// Add special VDEVs in compact form
	if v.pool.Cache != nil {
		output.WriteString("\nCache:\n  " + v.renderVDevCompact(v.pool.Cache))
	}
	if v.pool.Slog != nil {
		output.WriteString("\nSlog:\n  " + v.renderVDevCompact(v.pool.Slog))
	}

	return output.String()
}

// RenderDetailed shows expanded information for a specific VDEV
func (v *VDevView) RenderDetailed(vdevName string) string {
	var output strings.Builder

	// Find and render the selected VDEV in detail
	if vdev := v.findVDev(vdevName, v.pool.RootVDev); vdev != nil {
		output.WriteString(v.renderVDevDetailed(vdev))
	}
	if vdev := v.findVDev(vdevName, v.pool.Cache); vdev != nil {
		output.WriteString(v.renderVDevDetailed(vdev))
	}
	if vdev := v.findVDev(vdevName, v.pool.Slog); vdev != nil {
		output.WriteString(v.renderVDevDetailed(vdev))
	}

	return output.String()
}

// Helper to find a VDEV by name
func (v *VDevView) findVDev(name string, current *zfs.VDev) *zfs.VDev {
	if current == nil {
		return nil
	}
	if current.Name == name {
		return current
	}

	for _, child := range current.Children {
		if found := v.findVDev(name, child); found != nil {
			return found
		}
	}
	return nil
}

// renderDisks creates a formatted string representation of the physical disks
// in a VDev, including their status and resilvering progress if applicable
func (v *VDevView) renderDisks(disks []zfs.DiskInfo) string {
	var output string
	for _, disk := range disks {
		diskInfo := fmt.Sprintf("%s - %s (%s)\n", disk.Path, disk.Model, disk.Status)
		if disk.Resilvering {
			diskInfo += fmt.Sprintf("Resilvering: %.1f%%", disk.Progress)
		}
		output += diskInfo + "\n"
	}
	return output
}

func (v *VDevView) renderVDevCompact(vdev *zfs.VDev) string {
	var sb strings.Builder

	// Basic VDEV info
	status := renderStatus(vdev.Status)
	name := styles.VDevName.Render(vdev.Name)
	sb.WriteString(fmt.Sprintf("%s [%s]", name, status))

	if vdev.Path != "" {
		sb.WriteString(" (" + styles.VDevPath.Render(vdev.Path) + ")")
	}
	sb.WriteString("\n")

	// Render children in compact form
	for _, child := range vdev.Children {
		sb.WriteString("  " + v.renderVDevCompact(child))
	}

	return sb.String()
}

func (v *VDevView) renderVDevDetailed(vdev *zfs.VDev) string {
	var sb strings.Builder

	// Detailed VDEV info
	status := renderStatus(vdev.Status)
	name := styles.VDevName.Render(vdev.Name)
	sb.WriteString(fmt.Sprintf("%s [%s]\n", name, status))

	// Add additional details
	if vdev.Path != "" {
		sb.WriteString(fmt.Sprintf("  Path: %s\n", vdev.Path))
	}
	if vdev.State != "" {
		sb.WriteString(fmt.Sprintf("  State: %s\n", vdev.State))
	}

	// Render children in detail
	for _, child := range vdev.Children {
		sb.WriteString(v.renderVDevDetailed(child))
	}

	return sb.String()
}

func (v *VDevView) renderVDev(vdev *zfs.VDev) string {
	var s strings.Builder

	// Basic device info
	s.WriteString(styles.VDevName.Render(vdev.Name))
	if vdev.Path != "" {
		s.WriteString(" (" + styles.VDevPath.Render(vdev.Path) + ")")
	}

	return s.String()
}
