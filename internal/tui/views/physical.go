package views

import (
	"github.com/rivo/tview"
)

// PhysicalView creates the physical view showing vdev layout and disk health
func NewPhysicalView() *tview.Flex {
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

	// Sample physical topology
	poolNode := tview.NewTreeNode("tank").
		SetColor(tview.Styles.SecondaryTextColor)
	tree.GetRoot().AddChild(poolNode)

	// RAIDZ vdev
	raidzNode := tview.NewTreeNode("raidz2-0").
		SetColor(tview.Styles.TertiaryTextColor)
	poolNode.AddChild(raidzNode)

	// Physical disks
	disks := []string{"sda", "sdb", "sdc", "sdd"}
	for _, disk := range disks {
		diskNode := tview.NewTreeNode(disk + " [green](ONLINE)[white]").
			SetColor(tview.Styles.TertiaryTextColor)
		raidzNode.AddChild(diskNode)
	}

	// Cache and log devices
	cacheNode := tview.NewTreeNode("cache").
		SetColor(tview.Styles.TertiaryTextColor)
	poolNode.AddChild(cacheNode)

	cacheDisk := tview.NewTreeNode("nvme0n1 [green](ONLINE)[white]").
		SetColor(tview.Styles.TertiaryTextColor)
	cacheNode.AddChild(cacheDisk)

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
				{"Pool", "tank"},
				{"Status", "ONLINE"},
				{"Total Capacity", "4TB"},
				{"VDevs", "1 raidz2, 1 cache"},
				{"Resilver", "None in progress"},
				{"Scrub", "Last: 2025-01-01"},
			}
		} else if text == "raidz2-0" {
			properties = [][]string{
				{"VDev", "raidz2-0"},
				{"Type", "RAIDZ2"},
				{"Status", "ONLINE"},
				{"Devices", "4"},
				{"Capacity", "4TB"},
				{"Parity", "2 devices"},
				{"Resilver", "None"},
			}
		} else if text == "sda [green](ONLINE)[white]" {
			properties = [][]string{
				{"Device", "sda"},
				{"Status", "ONLINE"},
				{"Capacity", "1TB"},
				{"Read Errors", "0"},
				{"Write Errors", "0"},
				{"Checksum Errors", "0"},
				{"Temperature", "35°C"},
				{"Model", "WD Red 1TB"},
			}
		} else {
			properties = [][]string{
				{"Device", text},
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