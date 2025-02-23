package zfs

type VDevStatus string

const (
    VDevStatusOnline  VDevStatus = "ONLINE"
    VDevStatusDegraded VDevStatus = "DEGRADED"
    VDevStatusFaulted VDevStatus = "FAULTED"
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
}

func GetPools() ([]*Pool, error) {
    // TODO: Implement actual ZFS pool detection
    // For now, return mock data
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
                        Status: VDevStatusOnline,
                    },
                    {
                        Name:   "sdb",
                        Type:   "disk",
                        Status: VDevStatusOnline,
                    },
                },
            },
        },
    }, nil
}
