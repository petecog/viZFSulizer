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

type Pool struct {
	Name     string
	Status   VDevStatus
	RootVDev *VDev
	Cache    *VDev
	Slog     *VDev // Add SLOG VDev
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
