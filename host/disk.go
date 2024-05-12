package host

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"runtime"
)

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

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
		// 定义一个数组，用来忽略重复的分区，比如群晖
		var sli []string
		for _, partition := range partitions {
			usageStat, err := disk.Usage(partition.Mountpoint)
			if err != nil {
				continue
			}

			s := fmt.Sprintf("%v-%v", usageStat.Total, usageStat.Used)
			if contains(sli, s) {
				continue
			} else {
				sli = append(sli, s)
				total += usageStat.Total
				used += usageStat.Used
			}
		}
	}

	DiskTotal := total
	DiskUsed := used

	return DiskTotal, DiskUsed, nil
}
