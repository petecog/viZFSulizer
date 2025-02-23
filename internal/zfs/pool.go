package zfs

type VDevStatus string

const (
	VDevStatusOnline   VDevStatus = "ONLINE"
	VDevStatusDegraded VDevStatus = "DEGRADED"
	VDevStatusFaulted  VDevStatus = "FAULTED"
)

type VDev struct {
	Name     string
	Type     string
	Status   VDevStatus
	Children []*VDev
}

// GetWorstStatus returns the worst status among this VDev and its children
func (v *VDev) GetWorstStatus() VDevStatus {
	worst := v.Status

	for _, child := range v.Children {
		childStatus := child.GetWorstStatus()
		if childStatus == VDevStatusFaulted {
			return VDevStatusFaulted
		}
		if childStatus == VDevStatusDegraded && worst != VDevStatusFaulted {
			worst = VDevStatusDegraded
		}
	}
	return worst
}

type Pool struct {
	Name     string
	Status   VDevStatus
	RootVDev *VDev
	Cache    *VDev
	Slog     *VDev // Add SLOG VDev
}

// GetWorstStatus returns the worst status among this VDev and its children
func (p *Pool) GetWorstStatus() VDevStatus {
	worst := p.RootVDev.GetWorstStatus()

	// Check Cache devices
	if p.Cache != nil {
		if cacheStatus := p.Cache.GetWorstStatus(); isWorse(cacheStatus, worst) {
			worst = cacheStatus
		}
	}

	// Check SLOG devices
	if p.Slog != nil {
		if slogStatus := p.Slog.GetWorstStatus(); isWorse(slogStatus, worst) {
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

func GetPools() ([]*Pool, error) {
	return []*Pool{
		{
			Name:   "testpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				Name:   "testpool",
				Type:   "mirror",
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						Name:   "sda",
						Type:   "disk",
						Status: VDevStatusDegraded,
					},
					{
						Name:   "sdb",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
		},
		{
			Name:   "fastpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				Name:   "fastpool",
				Type:   "mirror",
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						Name:   "sda1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
					{
						Name:   "sdb1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
			Cache: &VDev{
				Name:   "cache",
				Type:   "cache",
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						Name:   "nvme0n1p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
			Slog: &VDev{
				Name:   "log",
				Type:   "mirror",
				Status: VDevStatusFaulted,
				Children: []*VDev{
					{
						Name:   "nvme1n1p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
					{
						Name:   "nvme1n2p1",
						Type:   "disk",
						Status: VDevStatusOnline,
					},
				},
			},
		},
	}, nil
}
