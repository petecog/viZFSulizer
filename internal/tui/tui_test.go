package tui

import (
	"testing"

	"github.com/petecog/vizfsulizer/internal/zfs"
)

func TestBasicFunctionality(t *testing.T) {
	// Add basic tests for your TUI components
	// Example:
	if 1+1 != 2 {
		t.Errorf("Basic math failed")
	}
}

func TestPoolNavigation(t *testing.T) {
	// Setup
	model := NewModel()
	pools := []*zfs.Pool{
		{Name: "pool1"},
		{Name: "pool2"},
		{Name: "pool3"},
	}
	model.pools = pools

	// Test initial state
	if model.selected != 0 {
		t.Errorf("Initial selection should be 0, got %d", model.selected)
	}

	// Test forward navigation
	model.selected = (model.selected + 1) % len(model.pools)
	if model.selected != 1 {
		t.Errorf("Forward navigation should be 1, got %d", model.selected)
	}

	// Test backward navigation
	model.selected = (model.selected - 1 + len(model.pools)) % len(model.pools)
	if model.selected != 0 {
		t.Errorf("Backward navigation should be 0, got %d", model.selected)
	}

	// Test wraparound
	model.selected = (model.selected - 1 + len(model.pools)) % len(model.pools)
	if model.selected != 2 {
		t.Errorf("Wraparound should be 2, got %d", model.selected)
	}
}
