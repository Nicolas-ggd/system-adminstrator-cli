package linegraph

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"os/exec"
)

// table header
var (
	tableHeaders = []string{"%CPU", "%MEM"}
)

// clear screen
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DrawTable draw table using external library
func DrawTable(cpuUsage, memUsage float64) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(tableHeaders)
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

	table.Append([]string{fmt.Sprintf("%.1f%%", cpuUsage), fmt.Sprintf("%.1f%%", memUsage)})

	return table
}
