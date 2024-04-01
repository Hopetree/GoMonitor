package host

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getNetworkSpeed(interval time.Duration) (uint64, uint64, error) {
	initialStats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}

	time.Sleep(interval)

	finalStats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}

	uploadSpeed := finalStats[0].BytesSent - initialStats[0].BytesSent
	downloadSpeed := finalStats[0].BytesRecv - initialStats[0].BytesRecv

	return uploadSpeed, downloadSpeed, nil
}

func formatSpeed(speed uint64) string {
	if speed < 1024 {
		return fmt.Sprintf("%dB", speed)
	} else if speed < 1024*1024 {
		return fmt.Sprintf("%.2fK", float64(speed)/1024)
	} else {
		return fmt.Sprintf("%.2fM", float64(speed)/(1024*1024))
	}
}

func getNetworkSpeedStr() (uploadSpeedStr, downloadSpeedStr string, err error) {
	interval := 1 * time.Second
	uploadSpeed, downloadSpeed, err := getNetworkSpeed(interval)

	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	uploadSpeedFormatted := formatSpeed(uploadSpeed)
	downloadSpeedFormatted := formatSpeed(downloadSpeed)
	return uploadSpeedFormatted, downloadSpeedFormatted, nil
}

func getConnections() (int, int, error) {
	var tcp, udp int
	cmd := exec.Command("sh", "-c", "expr $(ss -t | wc -l) - 1")

	// 获取命令输出
	tcpString, err := cmd.Output()
	if err == nil {
		tcpNum, err := strconv.Atoi(strings.TrimSpace(string(tcpString)))
		if err == nil {
			tcp = tcpNum
			if tcp < 0 {
				tcp = 0
			}
		}
	}

	cmd = exec.Command("sh", "-c", "expr $(ss -u | wc -l) - 1")

	// 获取命令输出
	udpString, err := cmd.Output()
	if err == nil {
		udpNum, err := strconv.Atoi(strings.TrimSpace(string(udpString)))
		if err == nil {
			udp = udpNum
			if udp < 0 {
				udp = 0
			}
		}
	}

	return tcp, udp, nil
}
