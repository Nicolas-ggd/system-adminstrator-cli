package monitor

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/parse"
	"github.com/fatih/color"
	pCPU "github.com/shirou/gopsutil/cpu"
	"log"
	"os"
	"strings"
)

var (
	processing = color.New(color.Bold, color.FgGreen)
)

// CPUStats hold the CPU times from /proc/stat
type CPUStats struct {
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	Iowait    uint64
	Irq       uint64
	Softirq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

// ReadCPUTasks function calculate CPU average usage percentage without idle
func ReadCPUTasks(cpuCount int) ([]CPUStats, error) {
	var stats []CPUStats
	var cpuStat CPUStats

	val, err := os.ReadFile(kernelDir)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(val), "\n")
	for _, line := range lines {
		for i := 0; i < cpuCount; i++ {
			if strings.HasPrefix(line, fmt.Sprintf("cpu%d ", i)) {
				fields := strings.Fields(line)
				if len(fields) < 11 {
					return nil, fmt.Errorf("unexpected format in /proc/stat")
				}

				cpuStat = CPUStats{
					User:      parse.ToUint64(fields[1]),
					Nice:      parse.ToUint64(fields[2]),
					System:    parse.ToUint64(fields[3]),
					Idle:      parse.ToUint64(fields[4]),
					Iowait:    parse.ToUint64(fields[5]),
					Irq:       parse.ToUint64(fields[6]),
					Softirq:   parse.ToUint64(fields[7]),
					Steal:     parse.ToUint64(fields[8]),
					Guest:     parse.ToUint64(fields[9]),
					GuestNice: parse.ToUint64(fields[10]),
				}

				stats = append(stats, cpuStat)
			}
		}
	}

	if len(stats) != cpuCount {
		return nil, fmt.Errorf("could not read CPU stats")
	}

	return stats, nil
}

func CalculateCPUUsage(start, end []CPUStats) ([]float64, error) {
	if len(start) != len(end) {
		return nil, fmt.Errorf("start and end slices must have the same length")
	}

	usage := make([]float64, len(start))

	for i := range start {
		startTotal := start[i].User + start[i].Nice + start[i].System + start[i].Idle + start[i].Iowait + start[i].Irq + start[i].Softirq + start[i].Steal
		endTotal := end[i].User + end[i].Nice + end[i].System + end[i].Idle + end[i].Iowait + end[i].Irq + end[i].Softirq + end[i].Steal

		totalDelta := endTotal - startTotal
		idleDelta := (end[i].Idle + end[i].Iowait) - (start[i].Idle + start[i].Iowait)

		if totalDelta == 0 {
			usage[i] = 0.0
		} else {
			usage[i] = 100 * float64(totalDelta-idleDelta) / float64(totalDelta)
		}
	}

	return usage, nil
}

// CountCPUCore function return number physical and logical cores count
func CountCPUCore() int {
	cpuCount, err := pCPU.Counts(true)
	if err != nil {
		log.Fatalf("Failed to count CPU %v", err.Error())
	}

	return cpuCount
}
