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

	// Details panel
	details := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Logical View[white]\nUse arrow keys to navigate the tree.\n\nSelect a pool or dataset to view properties.")

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
		ref := node.GetReference()
		
		var content string
		switch v := ref.(type) {
		case *zfs.Pool:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Name           %s\n"+
				"Status         %s\n"+
				"Health         %s\n"+
				"Size           %s\n"+
				"Allocated      %s\n"+
				"Free           %s\n"+
				"Dedup          %s\n"+
				"Last Scrub     %s",
				text, v.Name, v.Status, v.Health, v.Size, v.Allocated, v.Free, v.Dedup, v.LastScrub)
		case *zfs.Dataset:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Name           %s\n"+
				"Type           %s\n"+
				"Mountpoint     %s\n"+
				"Used           %s\n"+
				"Available      %s\n"+
				"Compression    %s\n"+
				"Dedup          %s\n"+
				"Snapshots      %d",
				text, v.Name, v.Type, v.Mountpoint, v.Used, v.Available, v.Compression, v.Dedup, len(v.Snapshots))
		case *zfs.Snapshot:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Name           %s\n"+
				"Type           Snapshot\n"+
				"Creation       %s\n"+
				"Used           %s\n"+
				"Referenced     %s",
				text, v.Name, v.Creation, v.Used, v.Referenced)
		default:
			content = fmt.Sprintf("[green]Selected: [white]%s\n\n"+
				"Type           Unknown\n"+
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

// newLogicalViewFallback creates a fallback view with hardcoded data
func newLogicalViewFallback() *tview.Flex {
	// Simple fallback implementation
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[red]Error loading ZFS config[white]\n\nFallback mode with sample data.\nCheck that config/zfs-config.yaml exists.")
	
	flex := tview.NewFlex().AddItem(text, 0, 1, true)
	return flex
}