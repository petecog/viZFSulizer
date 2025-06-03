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

	// Details panel with table
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)
	
	// Set up table headers
	headers := []string{"Property", "Value"}
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(header).
			SetTextColor(tview.Styles.SecondaryTextColor).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}
	
	// Default content
	table.SetCell(1, 0, tview.NewTableCell("Status").SetSelectable(false))
	table.SetCell(1, 1, tview.NewTableCell("Select a vdev or disk"))
	
	details := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Physical View[white]\nUse arrow keys to navigate the topology.")

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
		
		// Clear table and add data
		for row := 1; row < table.GetRowCount(); row++ {
			table.RemoveRow(row)
		}
		
		// Get data based on node reference
		var properties [][]string
		ref := node.GetReference()
		
		switch v := ref.(type) {
		case *zfs.Pool:
			properties = [][]string{
				{"Pool", v.Name},
				{"Status", v.Status},
				{"Health", v.Health},
				{"Total Capacity", v.Size},
				{"Allocated", v.Allocated},
				{"Free", v.Free},
				{"VDevs", fmt.Sprintf("%d", len(v.VDevs))},
				{"Last Scrub", v.LastScrub},
			}
		case *zfs.VDev:
			properties = [][]string{
				{"VDev", v.Name},
				{"Type", v.Type},
				{"Status", v.Status},
				{"Devices", fmt.Sprintf("%d", len(v.Devices))},
			}
		case *zfs.Device:
			properties = [][]string{
				{"Device", v.Name},
				{"Status", v.Status},
				{"Capacity", v.Capacity},
				{"Model", v.Model},
				{"Serial", v.Serial},
				{"Temperature", fmt.Sprintf("%d°C", v.Temperature)},
				{"Read Errors", fmt.Sprintf("%d", v.ReadErrors)},
				{"Write Errors", fmt.Sprintf("%d", v.WriteErrors)},
				{"Checksum Errors", fmt.Sprintf("%d", v.ChecksumErrors)},
			}
		default:
			properties = [][]string{
				{"Name", text},
				{"Status", "Select for details"},
			}
		}
		
		for i, prop := range properties {
			table.SetCell(i+1, 0, tview.NewTableCell(prop[0]).SetSelectable(false))
			table.SetCell(i+1, 1, tview.NewTableCell(prop[1]))
		}
		
		details.SetText("[green]Selected: [white]" + text + "\n[yellow]Use arrow keys to navigate.")
	})

	// Right panel with details text on top, table below
	rightPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(details, 3, 0, false).
		AddItem(table, 0, 1, false)

	// Split layout
	flex := tview.NewFlex().
		AddItem(tree, 0, 2, true).
		AddItem(rightPanel, 0, 3, false)

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