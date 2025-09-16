package host

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func getProcessCount() (int, error) {
	switch runtime.GOOS {
	case "linux":
		// Linux: 直接数 /proc 下的数字目录
		dir, err := os.ReadDir("/proc")
		if err != nil {
			return 0, err
		}
		count := 0
		for _, entry := range dir {
			if entry.IsDir() {
				if _, err := strconv.Atoi(entry.Name()); err == nil {
					count++
				}
			}
		}
		return count, nil
	default:
		processList, err := process.Processes()
		if err != nil {
			return 0, err
		}
		return len(processList), nil
	}
}

func getThreadCount() (int, error) {
	switch runtime.GOOS {
	case "linux":
		// Linux 优化路径
		dir, err := os.ReadDir("/proc")
		if err != nil {
			return 0, err
		}

		totalThreads := 0
		for _, entry := range dir {
			if !entry.IsDir() {
				continue
			}
			if _, err := strconv.Atoi(entry.Name()); err != nil {
				continue
			}

			taskDir := "/proc/" + entry.Name() + "/task"
			tasks, err := os.ReadDir(taskDir)
			if err != nil {
				continue
			}
			totalThreads += len(tasks)
		}
		return totalThreads, nil

	default:
		processes, err := process.Processes()
		if err != nil {
			return 0, err
		}

		totalThreads := 0
		for _, proc := range processes {
			count, err := proc.NumThreads()
			if err != nil {
				continue
			}
			totalThreads += int(count)
		}
		return totalThreads, nil
	}
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

	// 读取群晖系统信息
	filename := "/etc.defaults/VERSION"
	if info.Platform == "" && fileExists(filename) {
		osName, osVersion := readVersionFile(filename)
		systemVersion = fmt.Sprintf("%s-%s-%s-%s-%s", info.OS, info.KernelVersion,
			info.KernelArch, osName, osVersion)
	}

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

// 从文件中读取系统信息
func readVersionFile(filePath string) (string, string) {
	versionInfo := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 移除两边的空白字符
		line = strings.TrimSpace(line)
		// 跳过空行
		if line == "" {
			continue
		}
		// 分割行内容，获取键和值
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// 去除两边的引号
			value = strings.Trim(value, "\"")
			versionInfo[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return "", ""
	}

	osName := versionInfo["os_name"]
	productVersion := versionInfo["productversion"]

	return osName, productVersion
}

// 判断文件是否存在
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
