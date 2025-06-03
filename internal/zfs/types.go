package zfs

import "time"

// Device represents a physical storage device
type Device struct {
	Name           string `yaml:"name"`
	Status         string `yaml:"status"`
	Capacity       string `yaml:"capacity"`
	Model          string `yaml:"model"`
	Serial         string `yaml:"serial"`
	Temperature    int    `yaml:"temperature"`
	ReadErrors     int    `yaml:"read_errors"`
	WriteErrors    int    `yaml:"write_errors"`
	ChecksumErrors int    `yaml:"checksum_errors"`
}

// VDev represents a virtual device (mirror, raidz, etc.)
type VDev struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Status  string   `yaml:"status"`
	Devices []Device `yaml:"devices"`
}

// Snapshot represents a ZFS snapshot
type Snapshot struct {
	Name       string `yaml:"name"`
	Creation   string `yaml:"creation"`
	Used       string `yaml:"used"`
	Referenced string `yaml:"referenced"`
}

// Dataset represents a ZFS dataset/filesystem
type Dataset struct {
	Name        string     `yaml:"name"`
	Type        string     `yaml:"type"`
	Mountpoint  string     `yaml:"mountpoint"`
	Used        string     `yaml:"used"`
	Available   string     `yaml:"available"`
	Compression string     `yaml:"compression"`
	Dedup       string     `yaml:"dedup"`
	Snapshots   []Snapshot `yaml:"snapshots"`
}

// Pool represents a ZFS storage pool
type Pool struct {
	Name       string    `yaml:"name"`
	Status     string    `yaml:"status"`
	Health     string    `yaml:"health"`
	Size       string    `yaml:"size"`
	Allocated  string    `yaml:"allocated"`
	Free       string    `yaml:"free"`
	Dedup      string    `yaml:"dedup"`
	LastScrub  string    `yaml:"last_scrub"`
	VDevs      []VDev    `yaml:"vdevs"`
	Datasets   []Dataset `yaml:"datasets"`
}

// ZFSConfig represents the entire ZFS configuration
type ZFSConfig struct {
	Pools []Pool `yaml:"pools"`
}

// GetCreationTime parses the creation time string
func (s *Snapshot) GetCreationTime() (time.Time, error) {
	return time.Parse("2006-01-02", s.Creation)
}

// GetAllDevices returns all devices across all vdevs in a pool
func (p *Pool) GetAllDevices() []Device {
	var devices []Device
	for _, vdev := range p.VDevs {
		devices = append(devices, vdev.Devices...)
	}
	return devices
}

// GetDatasetByName finds a dataset by name within the pool
func (p *Pool) GetDatasetByName(name string) *Dataset {
	for i := range p.Datasets {
		if p.Datasets[i].Name == name {
			return &p.Datasets[i]
		}
	}
	return nil
}

// GetVDevByName finds a vdev by name within the pool
func (p *Pool) GetVDevByName(name string) *VDev {
	for i := range p.VDevs {
		if p.VDevs[i].Name == name {
			return &p.VDevs[i]
		}
	}
	return nil
}

// GetDeviceByName finds a device by name across all vdevs
func (p *Pool) GetDeviceByName(name string) *Device {
	for _, vdev := range p.VDevs {
		for i := range vdev.Devices {
			if vdev.Devices[i].Name == name {
				return &vdev.Devices[i]
			}
		}
	}
	return nil
}