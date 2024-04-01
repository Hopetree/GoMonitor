package host

import (
	"github.com/shirou/gopsutil/disk"
)

func getDiskUsage() (uint64, uint64, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return 0, 0, err
	}

	var total, used uint64

	for _, partition := range partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		total += usageStat.Total
		used += usageStat.Used
	}

	DiskTotal := total
	DiskUsed := used

	return DiskTotal, DiskUsed, nil
}
