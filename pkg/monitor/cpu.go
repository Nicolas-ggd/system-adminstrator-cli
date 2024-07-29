package monitor

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/parse"
	"github.com/fatih/color"
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
func ReadCPUTasks() (CPUStats, error) {
	val, err := os.ReadFile(kernelDir)
	if err != nil {
		return CPUStats{}, err
	}

	lines := strings.Split(string(val), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			if len(fields) < 11 {
				return CPUStats{}, fmt.Errorf("unexpected format in /proc/stat")
			}

			return CPUStats{
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
			}, nil
		}
	}

	return CPUStats{}, fmt.Errorf("cpu line not found in /proc/stat")
}

func CalculateCPUUsage(start, end CPUStats) float64 {
	startTotal := start.User + start.Nice + start.System + start.Idle + start.Iowait + start.Irq + start.Softirq + start.Steal + start.Guest + start.GuestNice
	endTotal := end.User + end.Nice + end.System + end.Idle + end.Iowait + end.Irq + end.Softirq + end.Steal + end.Guest + end.GuestNice

	totalDelta := endTotal - startTotal
	idleDelta := end.Idle - start.Idle

	if totalDelta == 0 {
		return 0.0
	}
	return (float64(totalDelta) - float64(idleDelta)) / float64(totalDelta) * 100.0
}
