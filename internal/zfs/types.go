package zfs

// VDevStatus represents the health status of a ZFS virtual device (VDev).
// It is implemented as a string type to represent different operational states.
type VDevStatus string

// ZFS VDev status constants represent the possible health states of a virtual device.
const (
	// VDevStatusOnline indicates the device is functioning normally
	VDevStatusOnline VDevStatus = "ONLINE"

	// VDevStatusDegraded indicates the device is operating with reduced functionality
	// or redundancy but is still able to handle I/O operations
	VDevStatusDegraded VDevStatus = "DEGRADED"

	// VDevStatusFaulted indicates the device has completely failed
	// or is not responding to I/O operations
	VDevStatusFaulted VDevStatus = "FAULTED"
)

// VDev represents a ZFS Virtual Device, which can be a physical device (disk),
// a logical device (mirror, raidz), or a special device (cache, log).
// VDevs form a tree structure where non-leaf nodes are logical devices
// containing other VDevs as children.
type VDev struct {
	// Name is the identifier for this VDev (e.g., "sda", "mirror-0")
	Name string

	// Type indicates the VDev's role (e.g., "disk", "mirror", "raidz")
	Type string

	// Status represents the current health state of this VDev
	Status VDevStatus

	// Children contains any child VDevs for logical devices
	// For example, a mirror VDev would have multiple disk VDevs as children
	Children []*VDev
}

// Pool represents a ZFS storage pool, which is the top-level container
// for data storage. A pool consists of one or more VDevs arranged in
// a specific configuration for redundancy and performance.
type Pool struct {
	// Name is the identifier for this pool
	Name string

	// Status represents the overall health state of the pool
	Status VDevStatus

	// RootVDev is the main storage VDev configuration
	// This typically contains the pool's primary storage devices
	RootVDev *VDev

	// Cache is an optional L2ARC cache device
	// Used to improve read performance with frequently accessed data
	Cache *VDev

	// Slog is an optional separate intent log device (ZIL)
	// Used to improve synchronous write performance
	Slog *VDev
}
