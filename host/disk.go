package host

import (
	"github.com/shirou/gopsutil/disk"
	"runtime"
)

func getDiskUsage() (uint64, uint64, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return 0, 0, err
	}

	var total, used uint64

	switch runtime.GOOS {
	case "darwin":
		for _, partition := range partitions {
			if partition.Mountpoint == "/" {
				usageStat, err := disk.Usage(partition.Mountpoint)
				if err != nil {
					continue
				}
				total += usageStat.Total
				used += usageStat.Used
			}
		}
	default:
		for _, partition := range partitions {
			usageStat, err := disk.Usage(partition.Mountpoint)
			if err != nil {
				continue
			}
			total += usageStat.Total
			used += usageStat.Used
		}

	}

	DiskTotal := total
	DiskUsed := used

	return DiskTotal, DiskUsed, nil
}
