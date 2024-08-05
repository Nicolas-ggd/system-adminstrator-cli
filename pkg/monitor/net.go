// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package monitor

import (
	"bufio"
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/parse"
	"os"
	"strings"
)

type NetMonitor struct {
	Interface     string
	ReceivedBytes int64
	TransmitBytes int64
}

type NetMonitorResponse struct {
	ReceivedBytes int64
	TotalReceived int64
	TransmitBytes int64
	TotalTransmit int64
}

func readNetDevLine(line string) (*NetMonitor, error) {
	fields := strings.Fields(line)

	if len(fields) < 17 {
		return nil, fmt.Errorf("invalid netdev line: %s", line)
	}

	return &NetMonitor{
		Interface:     strings.TrimSuffix(fields[0], ":"),
		ReceivedBytes: parse.ToInt64(fields[1]),
		TransmitBytes: parse.ToInt64(fields[9]),
	}, nil
}

func ReadNetDev() ([]NetMonitor, error) {
	var netStat []NetMonitor

	val, err := os.Open(kernelNet)
	if err != nil {
		return nil, fmt.Errorf("failed to open kernel netdev file: %v", err)
	}
	defer val.Close()

	scanner := bufio.NewScanner(val)
	// skip first to line
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		stat, err := readNetDevLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to read netdev line: %v", err)
		}
		netStat = append(netStat, *stat)
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return netStat, nil
}

func ReadNetUsage(start, end []NetMonitor) (*NetMonitorResponse, error) {
	var netStat NetMonitorResponse

	for i, initial := range start {
		final := end[i]
		if final.Interface == initial.Interface {
			rec := final.ReceivedBytes - initial.ReceivedBytes
			trn := final.TransmitBytes - initial.TransmitBytes
			netStat.TotalReceived += rec
			netStat.TotalTransmit += trn

			netStat.ReceivedBytes = rec
			netStat.TransmitBytes = trn
		}
	}

	return &netStat, nil
}
