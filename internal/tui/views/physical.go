package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/yourusername/vizfsulizer/internal/zfs"
)

// PhysicalView creates the physical view showing vdev layout and disk health
func NewPhysicalView() *tview.Flex {
	// Initialize ZFS simulator
	sim := zfs.NewSimulator()
	if err := sim.LoadDefaultConfig(); err != nil {
		// Fallback to hardcoded data if config fails
		return newPhysicalViewFallback()
	}
	// Tree view for physical topology
	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode("Physical Layout").SetColor(tview.Styles.PrimaryTextColor)).
		SetCurrentNode(tview.NewTreeNode("Physical Layout"))

	// Details panel
	details := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Physical View[white]\nUse arrow keys to navigate the topology.\n\nSelect a pool, vdev, or device to view properties.")

	// Load pools from simulator
	pools := sim.GetPools()
	for _, pool := range pools {
		poolNode := tview.NewTreeNode(pool.Name).
			SetColor(tview.Styles.SecondaryTextColor).
			SetReference(&pool)
		tree.GetRoot().AddChild(poolNode)

		// Add vdevs
		for _, vdev := range pool.VDevs {
			vdevNode := tview.NewTreeNode(vdev.Name).
				SetColor(tview.Styles.TertiaryTextColor).
				SetReference(&vdev)
			poolNode.AddChild(vdevNode)

			// Add devices
			for _, device := range vdev.Devices {
				statusColor := "[green]"
				if device.Status != "ONLINE" {
					statusColor = "[red]"
				}
				deviceNode := tview.NewTreeNode(fmt.Sprintf("%s %s(%s)[white]", device.Name, statusColor, device.Status)).
					SetColor(tview.Styles.TertiaryTextColor).
					SetReference(&device)
				vdevNode.AddChild(deviceNode)
			}
		}
	}

	// Handle selection
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		text := node.GetText()
		ref := node.GetReference()
		
		var content string
		switch v := ref.(type) {
		case *zfs.Pool:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Pool           %s\n"+
				"Status         %s\n"+
				"Health         %s\n"+
				"Capacity       %s\n"+
				"Allocated      %s\n"+
				"Free           %s\n"+
				"VDevs          %d\n"+
				"Last Scrub     %s",
				text, v.Name, v.Status, v.Health, v.Size, v.Allocated, v.Free, len(v.VDevs), v.LastScrub)
		case *zfs.VDev:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"VDev           %s\n"+
				"Type           %s\n"+
				"Status         %s\n"+
				"Devices        %d",
				text, v.Name, v.Type, v.Status, len(v.Devices))
		case *zfs.Device:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Device         %s\n"+
				"Status         %s\n"+
				"Capacity       %s\n"+
				"Model          %s\n"+
				"Serial         %s\n"+
				"Temperature    %d°C\n"+
				"Read Errors    %d\n"+
				"Write Errors   %d\n"+
				"Chksum Errors  %d",
				text, v.Name, v.Status, v.Capacity, v.Model, v.Serial, v.Temperature, v.ReadErrors, v.WriteErrors, v.ChecksumErrors)
		default:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Status         Select for details", text)
		}
		
		details.SetText(content)
	})

	// Split layout
	flex := tview.NewFlex().
		AddItem(tree, 0, 2, true).
		AddItem(details, 0, 3, false)

	return flex
}

// newPhysicalViewFallback creates a fallback view with hardcoded data
func newPhysicalViewFallback() *tview.Flex {
	// Simple fallback implementation
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[red]Error loading ZFS config[white]\n\nFallback mode with sample data.\nCheck that config/zfs-config.yaml exists.")
	
	flex := tview.NewFlex().AddItem(text, 0, 1, true)
	return flex
}