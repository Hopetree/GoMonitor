package host

import "github.com/shirou/gopsutil/mem"

func getMemoryInfo() (uint64, uint64, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}

	memoryTotal := memory.Total
	memoryUsed := memory.Used

	return memoryTotal, memoryUsed, nil
}

func getSwapInfo() (uint64, uint64, error) {
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		return 0, 0, err
	}

	swapTotal := swapInfo.Total
	swapUsed := swapInfo.Used

	return swapTotal, swapUsed, nil
}
