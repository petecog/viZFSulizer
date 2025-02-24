package status

import "github.com/petecog/vizfsulizer/internal/zfs"

// Analyzer provides methods for analyzing ZFS component health states.
// It implements recursive traversal of VDev trees to determine the overall
// health status of pools and their components.
type Analyzer struct{}

// GetVDevWorstStatus analyzes a VDev and its children for the worst status.
// It recursively traverses the VDev tree structure, comparing status values
// to find the most severe status condition.
//
// Parameters:
//   - vdev: The VDev to analyze, including all its child devices
//
// Returns:
//   - zfs.VDevStatus: The most severe status found in the VDev tree
//
// Example:
//
//	analyzer := &Analyzer{}
//	status := analyzer.GetVDevWorstStatus(pool.RootVDev)
func (an *Analyzer) GetVDevWorstStatus(vdev *zfs.VDev) zfs.VDevStatus {
	worst := vdev.Status

	for _, child := range vdev.Children {
		childStatus := an.GetVDevWorstStatus(child)
		if an.isWorse(childStatus, worst) {
			worst = childStatus
		}
	}
	return worst
}

// GetPoolWorstStatus analyzes all components of a pool for worst status.
// This includes the root VDev, cache devices (L2ARC), and log devices (ZIL).
// Each component tree is analyzed separately and the worst status is returned.
//
// Parameters:
//   - pool: The ZFS pool to analyze, including all its components
//
// Returns:
//   - zfs.VDevStatus: The most severe status found anywhere in the pool
//
// Example:
//
//	analyzer := &Analyzer{}
//	status := analyzer.GetPoolWorstStatus(myPool)
func (an *Analyzer) GetPoolWorstStatus(pool *zfs.Pool) zfs.VDevStatus {
	worst := an.GetVDevWorstStatus(pool.RootVDev)

	if pool.Cache != nil {
		if cacheStatus := an.GetVDevWorstStatus(pool.Cache); an.isWorse(cacheStatus, worst) {
			worst = cacheStatus
		}
	}

	if pool.Slog != nil {
		if slogStatus := an.GetVDevWorstStatus(pool.Slog); an.isWorse(slogStatus, worst) {
			worst = slogStatus
		}
	}

	return worst
}

// isWorse determines if status 'a' represents a worse condition than status 'b'.
// The severity order from worst to best is:
//  1. FAULTED  - Device is completely unavailable
//  2. DEGRADED - Device is operating with reduced functionality
//  3. ONLINE   - Device is functioning normally
//
// Parameters:
//   - a: First status to compare
//   - b: Second status to compare
//
// Returns:
//   - bool: true if status 'a' is worse than status 'b'
//
// Example:
//
//	analyzer := &Analyzer{}
//	isWorse := analyzer.isWorse(zfs.VDevStatusFaulted, zfs.VDevStatusOnline) // returns true
func (an *Analyzer) isWorse(a, b zfs.VDevStatus) bool {
	if a == zfs.VDevStatusFaulted {
		return true
	}
	if a == zfs.VDevStatusDegraded && b != zfs.VDevStatusFaulted {
		return true
	}
	return false
}
