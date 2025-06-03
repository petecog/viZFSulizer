package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/yourusername/vizfsulizer/internal/tui/views"
)

func Start() error {
	app := tview.NewApplication()

	// Create tabbed pages
	pages := tview.NewPages()
	
	// Add logical and physical views
	logicalView := views.NewLogicalView()
	physicalView := views.NewPhysicalView()
	
	pages.AddPage("Logical", logicalView, true, true)
	pages.AddPage("Physical", physicalView, true, false)

	// Create status bar
	statusBar := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[green]1[white]: Logical  [green]2[white]: Physical  [green]Tab[white]: Switch  [green]Ctrl+C[white]: Exit")

	// Main layout with pages on top, status bar on bottom
	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(statusBar, 1, 0, false)

	// Set up key bindings for tab switching
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case '1':
			pages.SwitchToPage("Logical")
			return nil
		case '2':
			pages.SwitchToPage("Physical")
			return nil
		}
		switch event.Key() {
		case tcell.KeyTab:
			// Toggle between pages
			current, _ := pages.GetFrontPage()
			if current == "Logical" {
				pages.SwitchToPage("Physical")
			} else {
				pages.SwitchToPage("Logical")
			}
			return nil
		}
		return event
	})

	return app.SetRoot(mainLayout, true).EnableMouse(true).Run()
}
