package host

import (
	"encoding/json"
	"github.com/Hopetree/GoMonitor/app"
)

type ServerInfo struct {
	Interval      int     `json:"interval"`
	Uptime        uint64  `json:"uptime"`
	System        string  `json:"system"`
	CPUCores      int     `json:"cpu_cores"`
	CPUModelName  string  `json:"cpu_model"`
	CPU           float64 `json:"cpu"`
	Load1         string  `json:"load_1"`
	Load5         string  `json:"load_5"`
	Load15        string  `json:"load_15"`
	MemoryTotal   uint64  `json:"memory_total"`
	MemoryUsed    uint64  `json:"memory_used"`
	SwapTotal     uint64  `json:"swap_total"`
	SwapUsed      uint64  `json:"swap_used"`
	DiskTotal     uint64  `json:"hdd_total"`
	DiskUsed      uint64  `json:"hdd_used"`
	NetWorkIn     string  `json:"network_in"`
	NetWorkOut    string  `json:"network_out"`
	Process       int     `json:"process"`
	Thread        int     `json:"thread"`
	TCP           int     `json:"tcp"`
	UDP           int     `json:"udp"`
	Version       string  `json:"version"`
	ClientVersion string  `json:"client_version"`
}

func GetHostInfo(interval int, cmdStr string) string {
	uptime, system, _ := getSystemInfo()

	cpuCores := getCPUCores()
	cpuModelName, _ := getCPUModelName()
	cpuUsed, _ := getTotalCPULoad()
	load1, load5, load15, _ := getLoad()

	memoryTotal, memoryUsed, _ := getMemoryInfo()
	swapTotal, swapUsed, _ := getSwapInfo()

	diskTotal, diskUsed, _ := getDiskUsage()

	upload, download, _ := getNetworkSpeedStr()

	processCount, threadCount, _ := getProcessAndThread()
	tcpCount, udpCount, _ := getConnections()

	version := getVersion(cmdStr)

	serverInfo := ServerInfo{
		Interval:      interval,
		Uptime:        uptime,
		System:        system,
		CPUCores:      cpuCores,
		CPUModelName:  cpuModelName,
		CPU:           cpuUsed,
		Load1:         load1,
		Load5:         load5,
		Load15:        load15,
		MemoryTotal:   memoryTotal,
		MemoryUsed:    memoryUsed,
		SwapTotal:     swapTotal,
		SwapUsed:      swapUsed,
		DiskTotal:     diskTotal,
		DiskUsed:      diskUsed,
		NetWorkIn:     download,
		NetWorkOut:    upload,
		Process:       processCount,
		Thread:        threadCount,
		TCP:           tcpCount,
		UDP:           udpCount,
		Version:       version,
		ClientVersion: app.RuntimeVersion,
	}

	jsonData, _ := json.Marshal(&serverInfo)
	return string(jsonData)
}
