package monitor

import (
	psMem "github.com/shirou/gopsutil/mem"
	"log"
)

type MemoryInfo struct {
	Total       int64
	Used        int64
	UsedPercent float64
}

func NewMemoryInfo() *psMem.VirtualMemoryStat {
	mainMemory, err := psMem.VirtualMemory()
	if err != nil {
		log.Fatal("Failed to get virtual memory", err)
	}

	return mainMemory
}
