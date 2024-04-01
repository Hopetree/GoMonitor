package host

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
)

func getProcessCount() (int, error) {
	processes, err := process.Processes()
	if err != nil {
		return 0, err
	}

	return len(processes), nil
}

func getThreadCount() (int, error) {
	processes, err := process.Processes()
	if err != nil {
		return 0, err
	}

	var totalThreads int
	for _, proc := range processes {
		threadCount, err := proc.NumThreads()
		if err != nil {
			continue
		}
		totalThreads += int(threadCount)
	}

	return totalThreads, nil
}

func getProcessAndThread() (int, int, error) {
	processCount, err := getProcessCount()
	if err != nil {
		return 0, 0, err
	}
	threadCount, err := getThreadCount()
	if err != nil {
		return 0, 0, err
	}
	return processCount, threadCount, nil
}

func getSystemInfo() (uint64, string, error) {
	info, err := host.Info()
	if err != nil {
		return 0, "", err
	}

	uptime := info.Uptime
	systemVersion := fmt.Sprintf("%s-%s-%s-%s-%s", info.OS, info.KernelVersion,
		info.KernelArch, info.Platform, info.PlatformVersion)
	return uptime, systemVersion, nil
}
