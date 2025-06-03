package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/yourusername/vizfsulizer/internal/zfs"
)

// LogicalView creates the logical view showing pools and datasets hierarchy
func NewLogicalView() *tview.Flex {
	// Initialize ZFS simulator
	sim := zfs.NewSimulator()
	if err := sim.LoadDefaultConfig(); err != nil {
		// Fallback to hardcoded data if config fails
		return newLogicalViewFallback()
	}
	// Tree view for ZFS hierarchy
	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode("ZFS Pools").SetColor(tview.Styles.PrimaryTextColor)).
		SetCurrentNode(tview.NewTreeNode("ZFS Pools"))

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
	table.SetCell(1, 1, tview.NewTableCell("Select a pool or dataset"))
	
	details := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Logical View[white]\nUse arrow keys to navigate the tree.")

	// Load pools from simulator
	pools := sim.GetPools()
	for _, pool := range pools {
		poolNode := tview.NewTreeNode(pool.Name).
			SetColor(tview.Styles.SecondaryTextColor).
			SetReference(&pool)
		tree.GetRoot().AddChild(poolNode)

		// Add datasets
		for _, dataset := range pool.Datasets {
			datasetNode := tview.NewTreeNode(dataset.Name).
				SetColor(tview.Styles.TertiaryTextColor).
				SetReference(&dataset)
			poolNode.AddChild(datasetNode)

			// Add snapshots
			for _, snapshot := range dataset.Snapshots {
				snapshotNode := tview.NewTreeNode(snapshot.Name).
					SetColor(tview.Styles.TertiaryTextColor).
					SetReference(&snapshot)
				datasetNode.AddChild(snapshotNode)
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
				{"Name", v.Name},
				{"Status", v.Status},
				{"Health", v.Health},
				{"Size", v.Size},
				{"Allocated", v.Allocated},
				{"Free", v.Free},
				{"Dedup", v.Dedup},
				{"Last Scrub", v.LastScrub},
			}
		case *zfs.Dataset:
			properties = [][]string{
				{"Name", v.Name},
				{"Type", v.Type},
				{"Mountpoint", v.Mountpoint},
				{"Used", v.Used},
				{"Available", v.Available},
				{"Compression", v.Compression},
				{"Dedup", v.Dedup},
				{"Snapshots", fmt.Sprintf("%d", len(v.Snapshots))},
			}
		case *zfs.Snapshot:
			properties = [][]string{
				{"Name", v.Name},
				{"Type", "Snapshot"},
				{"Creation", v.Creation},
				{"Used", v.Used},
				{"Referenced", v.Referenced},
			}
		default:
			properties = [][]string{
				{"Name", text},
				{"Type", "Unknown"},
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

// newLogicalViewFallback creates a fallback view with hardcoded data
func newLogicalViewFallback() *tview.Flex {
	// Simple fallback implementation
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[red]Error loading ZFS config[white]\n\nFallback mode with sample data.\nCheck that config/zfs-config.yaml exists.")
	
	flex := tview.NewFlex().AddItem(text, 0, 1, true)
	return flex
}