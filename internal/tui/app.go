package tui

import (
	"github.com/rivo/tview"
)

func Start() error {
	app := tview.NewApplication()

	// Create main layout with tree on left, details on right
	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode("ZFS Pools").SetColor(tview.Styles.PrimaryTextColor)).
		SetCurrentNode(tview.NewTreeNode("ZFS Pools"))

	details := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow]Select a pool or dataset to view details[white]")

	// Split layout: 40% tree, 60% details
	flex := tview.NewFlex().
		AddItem(tree, 0, 2, true).   // tree gets 40% (2/5)
		AddItem(details, 0, 3, false) // details gets 60% (3/5)

	// Add sample data to tree
	poolNode := tview.NewTreeNode("tank").
		SetColor(tview.Styles.SecondaryTextColor)
	tree.GetRoot().AddChild(poolNode)

	datasetNode := tview.NewTreeNode("tank/data").
		SetColor(tview.Styles.TertiaryTextColor)
	poolNode.AddChild(datasetNode)

	// Handle tree selection
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		text := node.GetText()
		details.SetText("[green]Selected: [white]" + text + "\n\n[yellow]Details will go here...")
	})

	return app.SetRoot(flex, true).EnableMouse(true).Run()
}
