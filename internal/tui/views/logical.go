package views

import (
	"github.com/rivo/tview"
)

// LogicalView creates the logical view showing pools and datasets hierarchy
func NewLogicalView() *tview.Flex {
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

	// Sample data
	poolNode := tview.NewTreeNode("tank").
		SetColor(tview.Styles.SecondaryTextColor)
	tree.GetRoot().AddChild(poolNode)

	datasetNode := tview.NewTreeNode("tank/data").
		SetColor(tview.Styles.TertiaryTextColor)
	poolNode.AddChild(datasetNode)

	snapshotNode := tview.NewTreeNode("tank/data@snapshot1").
		SetColor(tview.Styles.TertiaryTextColor)
	datasetNode.AddChild(snapshotNode)

	// Handle selection
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		text := node.GetText()
		
		// Clear table and add data
		for row := 1; row < table.GetRowCount(); row++ {
			table.RemoveRow(row)
		}
		
		// Sample data based on selection
		var properties [][]string
		if text == "tank" {
			properties = [][]string{
				{"Name", "tank"},
				{"Status", "ONLINE"},
				{"Size", "1.2TB"},
				{"Used", "800GB (67%)"},
				{"Available", "400GB"},
				{"Health", "Healthy"},
				{"Dedup", "1.00x"},
				{"Compression", "lz4"},
			}
		} else if text == "tank/data" {
			properties = [][]string{
				{"Name", "tank/data"},
				{"Type", "Dataset"},
				{"Used", "600GB"},
				{"Available", "600GB"},
				{"Compression", "lz4"},
				{"Snapshots", "3"},
				{"Mountpoint", "/tank/data"},
			}
		} else {
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