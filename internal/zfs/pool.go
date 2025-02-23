package zfs

func GetPools() ([]*Pool, error) {
	return []*Pool{
		{
			Name:   "testpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				Name:   "testpool",
				Type:   VDevTypeMirror,
				Status: VDevStatusOnline,
				Children: []*VDev{
					{
						Name:   "sda",
						Type:   VDevTypeDisk,
						Status: VDevStatusDegraded,
						State:  "DEGRADED",
					},
					{
						Name:   "sdb",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
				},
			},
		},
		{
			Name:   "fastpool",
			Status: VDevStatusOnline,
			RootVDev: &VDev{
				Name:   "fastpool",
				Type:   VDevTypeMirror,
				Status: VDevStatusOnline,
				State:  "AVAIL",
				Children: []*VDev{
					{
						Name:   "sda1",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
					{
						Name:   "sdb1",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
				},
			},
			Cache: &VDev{
				Name:   "cache",
				Type:   VDevTypeCache,
				Status: VDevStatusOnline,
				State:  "AVAIL",
				Children: []*VDev{
					{
						Name:   "nvme0n1p1",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
				},
			},
			Slog: &VDev{
				Name:   "log",
				Type:   VDevTypeMirror,
				Status: VDevStatusFaulted,
				State:  "FAULTED",
				Children: []*VDev{
					{
						Name:   "nvme1n1p1",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
					{
						Name:   "nvme1n2p1",
						Type:   VDevTypeDisk,
						Status: VDevStatusOnline,
						State:  "AVAIL",
					},
				},
			},
		},
	}, nil
}
