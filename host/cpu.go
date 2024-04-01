package host

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"runtime"
	"time"
)

func getCPUModelName() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return "", err
	}

	var cpuModel string

	for _, info := range cpuInfo {
		cpuModel = info.ModelName
		break
	}
	return cpuModel, nil
}

func getTotalCPULoad() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, err
	}

	totalLoad := 0.0
	for _, eachLoad := range percentages {
		totalLoad += eachLoad
	}

	averageLoad := totalLoad / float64(len(percentages))
	return averageLoad, nil
}

func getLoad() (string, string, string, error) {
	loadAvg, err := load.Avg()
	if err != nil {
		return "", "", "", err
	}

	load1 := fmt.Sprintf("%.2f", loadAvg.Load1)
	load5 := fmt.Sprintf("%.2f", loadAvg.Load5)
	load15 := fmt.Sprintf("%.2f", loadAvg.Load15)

	return load1, load5, load15, nil
}

func getCPUCores() int {
	return runtime.NumCPU()
}
