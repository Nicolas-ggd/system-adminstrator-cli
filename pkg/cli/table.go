// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package cli

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/monitor"
	"github.com/olekukonko/tablewriter"
	"os"
	"os/exec"
)

// ClearScreen clear previous screen
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DrawTable draw table using external library
func DrawTable(cpuUsage []float64, mem *monitor.MemStatResponse) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	cpuHeader := "%Cpu(s):"
	cpuValues := ""
	for key, val := range cpuUsage {
		cpuValues += fmt.Sprintf("CPU%d %.1f%%, ", key, val)
	}

	// Remove the trailing comma and space
	if len(cpuValues) > 0 {
		cpuValues = cpuValues[:len(cpuValues)-2]
	}

	table.Append([]string{fmt.Sprintf("%s %s", cpuHeader, cpuValues)})
	table.Append([]string{fmt.Sprintf("Mib Mem: %.1f MemTotal, %.1f MemFree, %.1f MemUsed", mem.MemoryTotal, mem.MemFree, mem.MemoryUsed)})
	table.Append([]string{fmt.Sprintf("Mib Swap: %.1f SwapTotal, %.1f SwapFree, %.1f SwapUsed", mem.SwapTotal, mem.SwapFree, mem.SwapUsed)})

	return table
}
