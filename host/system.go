package host

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"os"
	"strconv"
	"strings"
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

	if isLXCContainer() {
		uptime, _ = getLXCContainerUptime()
	}

	return uptime, systemVersion, nil
}

// 判断当前是否为lxc的虚拟机，这种虚拟机获取uptime的时候会拿到宿主机的，所以要单独处理
func isLXCContainer() bool {
	data, err := os.ReadFile("/proc/1/environ")
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "container=lxc")
}

// 获取LXC容器的启动时间
func getLXCContainerUptime() (uint64, error) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0, err
	}
	uptimeStr := strings.Fields(string(data))[0]
	uptimeSeconds, err := strconv.ParseFloat(uptimeStr, 64)
	if err != nil {
		return 0, err
	}
	return uint64(uptimeSeconds), nil
}
