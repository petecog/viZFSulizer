package zfs

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Simulator loads and provides ZFS configuration data
type Simulator struct {
	config ZFSConfig
}

// NewSimulator creates a new ZFS simulator
func NewSimulator() *Simulator {
	return &Simulator{}
}

// LoadConfig loads ZFS configuration from a YAML file
func (s *Simulator) LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(data, &s.config)
	if err != nil {
		return fmt.Errorf("failed to parse YAML config: %w", err)
	}

	return nil
}

// LoadDefaultConfig loads the default config from the project
func (s *Simulator) LoadDefaultConfig() error {
	// Try to find the config file relative to the project root
	configPaths := []string{
		"config/zfs-config.yaml",
		"../config/zfs-config.yaml",
		"../../config/zfs-config.yaml",
	}
	
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			return s.LoadConfig(path)
		}
	}
	
	return fmt.Errorf("could not find zfs-config.yaml in any of the expected locations: %v", configPaths)
}

// GetPools returns all configured pools
func (s *Simulator) GetPools() []Pool {
	return s.config.Pools
}

// GetPool returns a specific pool by name
func (s *Simulator) GetPool(name string) *Pool {
	for i := range s.config.Pools {
		if s.config.Pools[i].Name == name {
			return &s.config.Pools[i]
		}
	}
	return nil
}

// GetPoolNames returns a list of all pool names
func (s *Simulator) GetPoolNames() []string {
	names := make([]string, len(s.config.Pools))
	for i, pool := range s.config.Pools {
		names[i] = pool.Name
	}
	return names
}

// GetDatasetNames returns all dataset names for a given pool
func (s *Simulator) GetDatasetNames(poolName string) []string {
	pool := s.GetPool(poolName)
	if pool == nil {
		return nil
	}
	
	names := make([]string, len(pool.Datasets))
	for i, dataset := range pool.Datasets {
		names[i] = dataset.Name
	}
	return names
}

// GetVDevNames returns all vdev names for a given pool
func (s *Simulator) GetVDevNames(poolName string) []string {
	pool := s.GetPool(poolName)
	if pool == nil {
		return nil
	}
	
	names := make([]string, len(pool.VDevs))
	for i, vdev := range pool.VDevs {
		names[i] = vdev.Name
	}
	return names
}

// GetDeviceNames returns all device names for a given pool
func (s *Simulator) GetDeviceNames(poolName string) []string {
	pool := s.GetPool(poolName)
	if pool == nil {
		return nil
	}
	
	var names []string
	for _, vdev := range pool.VDevs {
		for _, device := range vdev.Devices {
			names = append(names, device.Name)
		}
	}
	return names
}

// UpdateDeviceStatus updates the status of a specific device
func (s *Simulator) UpdateDeviceStatus(poolName, deviceName, status string) error {
	pool := s.GetPool(poolName)
	if pool == nil {
		return fmt.Errorf("pool %s not found", poolName)
	}
	
	device := pool.GetDeviceByName(deviceName)
	if device == nil {
		return fmt.Errorf("device %s not found in pool %s", deviceName, poolName)
	}
	
	device.Status = status
	return nil
}

// UpdatePoolStatus updates the status of a pool
func (s *Simulator) UpdatePoolStatus(poolName, status string) error {
	pool := s.GetPool(poolName)
	if pool == nil {
		return fmt.Errorf("pool %s not found", poolName)
	}
	
	pool.Status = status
	return nil
}

// SaveConfig saves the current configuration to a file
func (s *Simulator) SaveConfig(configPath string) error {
	// Ensure directory exists
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	data, err := yaml.Marshal(&s.config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to YAML: %w", err)
	}
	
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

// GetConfig returns the raw configuration (useful for debugging)
func (s *Simulator) GetConfig() *ZFSConfig {
	return &s.config
}