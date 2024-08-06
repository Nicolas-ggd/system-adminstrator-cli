// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package monitor

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/parse"
	"log"
	"os"
	"strings"
)

type MemStats struct {
	MemTotal  int64
	MemFree   int64
	SwapTotal int64
	SwapFree  int64
	Buffers   int64
	Cached    int64
}

type MemStatResponse struct {
	MemoryTotal      float64
	MemoryUsed       float64
	MemoryPercentage float64
	MemFree          float64
	SwapTotal        float64
	SwapUsed         float64
	SwapPercentage   float64
	SwapFree         float64
}

func ReadMemUsage() (*MemStatResponse, error) {
	var mem MemStats

	val, err := os.ReadFile(kernelMem)
	if err != nil {
		return nil, fmt.Errorf("failed to read mem usage: %v\n", err)
	}

	lines := strings.Split(string(val), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 1 && fields[0] == "MemTotal:" {
			mem.MemTotal = parse.ToInt64(fields[1])
		} else if len(fields) > 1 && fields[0] == "SwapTotal:" {
			mem.SwapTotal = parse.ToInt64(fields[1])
		} else if len(fields) > 1 && fields[0] == "Buffers:" {
			mem.Buffers = parse.ToInt64(fields[1])
		} else if len(fields) > 1 && fields[0] == "Cached:" {
			mem.Cached = parse.ToInt64(fields[1])
		} else if len(fields) > 1 && fields[0] == "MemFree:" {
			mem.MemFree = parse.ToInt64(fields[1])
		} else if len(fields) > 1 && fields[0] == "SwapFree:" {
			mem.SwapFree = parse.ToInt64(fields[1])
		}

	}

	usedMem := mem.MemTotal - mem.MemFree - mem.Buffers - mem.Cached
	percentageMem := float64(usedMem) / float64(mem.MemTotal) * 100

	usedSwap := mem.SwapTotal - getSwapFree()
	percentageSwap := float64(usedSwap) / float64(mem.SwapTotal) * 100

	memResp := &MemStatResponse{
		MemoryTotal:      parse.KBToMib(mem.MemTotal),
		MemoryUsed:       parse.KBToMib(usedMem),
		MemoryPercentage: percentageMem,
		MemFree:          parse.KBToMib(mem.MemFree),
		SwapTotal:        parse.KBToMib(mem.SwapTotal),
		SwapUsed:         parse.KBToMib(usedSwap),
		SwapPercentage:   percentageSwap,
		SwapFree:         parse.KBToMib(mem.SwapFree),
	}

	return memResp, nil
}

func getSwapFree() int64 {
	val, err := os.ReadFile(kernelMem)
	if err != nil {
		log.Fatalf("failed to read mem usage: %v\n", err)
	}

	lines := strings.Split(string(val), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 1 && fields[0] == "SwapFree:" {
			return parse.ToInt64(fields[1])
		}
	}

	return 0
}
