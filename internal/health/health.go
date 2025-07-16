package health

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsagePercent float64 `json:"cpu_percent"`
	CPUTemp         float64 `json:"cpu_temp,omitempty"`
	RAMUsedMB       uint64  `json:"ram_used_mb"`
	RAMTotalMB      uint64  `json:"ram_total_mb"`
	DiskUsedMB      uint64  `json:"disk_used_gb"`
	DiskTotalMB     uint64  `json:"disk_total_gb"`
}

func GetSystemStats() (*SystemStats, error) {
	// CPU usage (per 1s sample)
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil || len(cpuPercent) == 0 {
		return nil, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	// Cpu temperature
	var CpuTemp float64
	temps, err := host.SensorsTemperatures()
	if err == nil {
		for _, t := range temps {
			if t.SensorKey == "Package id 0" || t.SensorKey == "CPU Temperature" {
				CpuTemp = t.Temperature
				break
			}
		}
	}

	// RAM usage
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get RAM usage: %w", err)
	}

	// Disk usage (root filesystem)
	du, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage: %w", err)
	}

	stats := &SystemStats{
		CPUUsagePercent: cpuPercent[0],
		CPUTemp:         CpuTemp,
		RAMUsedMB:       vm.Used / 1024 / 1024,
		RAMTotalMB:      vm.Total / 1024 / 1024,
		DiskUsedMB:      du.Used / 1024 / 1024,
		DiskTotalMB:     du.Total / 1024 / 1024,
	}

	return stats, nil
}
