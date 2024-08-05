// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package cli

import (
	"fmt"
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
func DrawTable(cpuUsage []float64) *tablewriter.Table {
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

	for key, val := range cpuUsage {
		table.Append([]string{fmt.Sprintf("CPU%d %.1f%%", key, val)})
	}

	return table
}
