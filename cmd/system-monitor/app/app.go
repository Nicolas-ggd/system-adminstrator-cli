package app

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/linegraph"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/monitor"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/mem"
	"os"
	"runtime"
	"time"
)

var (
	processing = color.New(color.Bold, color.FgGreen)
	invalid    = color.New(color.Bold, color.FgRed)
)

func Run() {
	systemOs := detectOS()

	arg := os.Args[1:]

	processing.Println("➜ system-monitor starting...")
	switch arg[0] {
	case "run":
		switch systemOs {
		case "linux":
			startStats, err := monitor.ReadCPUTasks()
			if err != nil {
				invalid.Printf("Error reading CPU stats, failed to %s\n", err.Error())
				invalid.Printf("Invalid OS system, your current OS is: %s\n", systemOs)
				os.Exit(0)
			}

			for {
				time.Sleep(2 * time.Second)
				endStats, err := monitor.ReadCPUTasks()
				if err != nil {
					invalid.Printf("Error reading CPU stats: %s\n", err.Error())
				}
				v, _ := mem.VirtualMemory()
				fmt.Println(v.Free)

				cpuUsage := monitor.CalculateCPUUsage(startStats, endStats)

				// clear screen
				linegraph.ClearScreen()

				// draw table
				table := linegraph.DrawTable(cpuUsage, v.UsedPercent)

				// render table
				table.Render()
			}
		default:
			os.Exit(0)
		}
	case "info":
		monitor.CpuLogger()
		os.Exit(0)
	default:
		help()
		os.Exit(0)
	}
}

func help() {
	c := color.New(color.FgRed)
	text := "Please provide a command to access system administrators."
	help := `
Usage
	$ system-monitor <run command>

Examples
	$ system-monitor run 
	$ system-monitor info 
	`
	_, err := c.Printf("%s %s\n", text, help)
	if err != nil {
		return
	}

}

// detectOS is used to detect operating system and the target architecture of a running program.
func detectOS() string {
	osSystem := runtime.GOOS

	switch osSystem {
	case "linux":
		processing.Println("➜ OS: Linux")
		processing.Printf("➜ Architecture: %s\n", osSystem)
	default:
		processing.Printf("%s.\n", osSystem)
	}

	return osSystem
}
