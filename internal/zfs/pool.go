package zfs

// GetPools returns a list of ZFS storage pools and their current state.
// Currently provides mock data for development and testing purposes.
// TODO: In production, this will be replaced with actual ZFS command execution.
//
// Returns:
//   - []*Pool: Slice of pointers to Pool structures containing pool configurations
//   - error: Error if pool data cannot be retrieved (currently always nil)
//
// Example:
//
//	pools, err := GetPools()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, pool := range pools {
//	    fmt.Printf("Pool: %s Status: %s\n", pool.Name, pool.Status)
//	}
func GetPools() ([]*Pool, error) {
	// Mock data representing two ZFS pools with different configurations
	return []*Pool{
		{
			// Basic mirrored pool configuration
			Name:   "testpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				// Root VDev represents the main storage configuration
				Name:   "testpool",
				Type:   "mirror", // Mirror provides 2-way redundancy
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						// First disk in mirror is degraded
						Name:   "sda",
						Type:   "disk",
						Status: VDevStatusDegraded,
					},
					{
						// Second disk in mirror is healthy
						Name:   "sdb",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
		},
		{
			// Advanced pool configuration with cache and log devices
			Name:   "fastpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				// Main storage configuration using mirrored disks
				Name:   "fastpool",
				Type:   "mirror",
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						// First disk partition in mirror
						Name:   "sda1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
					{
						// Second disk partition in mirror
						Name:   "sdb1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
			Cache: &VDev{
				// L2ARC cache device for read performance
				Name:   "cache",
				Type:   "cache",
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						// Using NVMe drive for better cache performance
						Name:   "nvme0n1p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
			Slog: &VDev{
				// ZFS Intent Log (ZIL) for sync write performance
				// Currently in faulted state despite healthy devices
				Name:   "log",
				Type:   "mirror",
				Status: VDevStatusFaulted,
				Children: []*VDev{
					{
						// First NVMe partition for log mirror
						Name:   "nvme1n1p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
					{
						// Second NVMe partition for log mirror
						Name:   "nvme1n2p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
		},
	}, nil
}
