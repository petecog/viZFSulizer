package zfs

// VDevType represents the type of a virtual device in a ZFS pool
type VDevType string

// VDevStatus represents the operational status of a virtual device
type VDevStatus string

// VDev type constants
const (
	VDevTypeMirror VDevType = "mirror"
	VDevTypeRaidz1 VDevType = "raidz1"
	VDevTypeRaidz2 VDevType = "raidz2"
	VDevTypeRaidz3 VDevType = "raidz3"
	VDevTypeSpare  VDevType = "spare"
	VDevTypeCache  VDevType = "cache"
	VDevTypeLog    VDevType = "log"
	VDevTypeSingle VDevType = "single"
	VDevTypeDisk   VDevType = "disk"
)

// Status constants for virtual devices
const (
	VDevStatusOnline   VDevStatus = "ONLINE"
	VDevStatusDegraded VDevStatus = "DEGRADED"
	VDevStatusFaulted  VDevStatus = "FAULTED"
)

// DiskInfo contains detailed information about a physical disk device
type DiskInfo struct {
	Name        string
	Path        string
	Size        int64
	Model       string
	Serial      string
	Status      VDevStatus
	ReadOps     int64
	WriteOps    int64
	Resilvering bool
	Progress    float64
}

// VDev represents a virtual device in a ZFS pool
type VDev struct {
	Type          VDevType
	Name          string
	Status        VDevStatus
	Path          string // Add Path field for device location
	Children      []*VDev
	Disks         []DiskInfo
	Redundancy    int
	UsedCapacity  uint64
	TotalCapacity uint64
	State         string
}

// GetWorstStatus returns the worst status among the VDev and its children
func (v *VDev) GetWorstStatus() VDevStatus {
	worst := v.Status

	for _, child := range v.Children {
		childStatus := child.GetWorstStatus()
		if isWorse(childStatus, worst) {
			worst = childStatus
		}
	}

	// If no children, but disks present, return worst disk status.
	if len(v.Children) == 0 && len(v.Disks) > 0 {
		for _, disk := range v.Disks {
			if isWorse(disk.Status, worst) {
				worst = disk.Status
			}
		}
	}

	return worst
}

// IsRedundant returns true if the VDEV type provides redundancy
func (v *VDev) IsRedundant() bool {
	return v.Type == VDevTypeMirror ||
		v.Type == VDevTypeRaidz1 ||
		v.Type == VDevTypeRaidz2 ||
		v.Type == VDevTypeRaidz3
}

type Pool struct {
	Name     string
	Status   VDevStatus
	RootVDev *VDev
	Cache    *VDev
	Slog     *VDev // Add SLOG VDev
}

// GetWorstStatus returns the worst status among the pool and its components (RootVDev, Cache, Slog)
func (p *Pool) GetWorstStatus() VDevStatus {
	worst := p.Status

	if p.RootVDev != nil {
		rootStatus := p.RootVDev.GetWorstStatus()
		if isWorse(rootStatus, worst) {
			worst = rootStatus
		}
	}

	if p.Cache != nil {
		cacheStatus := p.Cache.GetWorstStatus()
		if isWorse(cacheStatus, worst) {
			worst = cacheStatus
		}
	}

	if p.Slog != nil {
		slogStatus := p.Slog.GetWorstStatus()
		if isWorse(slogStatus, worst) {
			worst = slogStatus
		}
	}

	return worst
}

// isWorse returns true if status a is worse than status b
func isWorse(a, b VDevStatus) bool {
	if a == VDevStatusFaulted {
		return true
	}
	if a == VDevStatusDegraded && b != VDevStatusFaulted {
		return true
	}
	return false
}
