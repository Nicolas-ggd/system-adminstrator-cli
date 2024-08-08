package cli

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func DrawTableTop(headers []string, data [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding(" ") // pad with tabs
	table.SetNoWhiteSpace(false)
	table.AppendBulk(data)

	return table
}
