package app

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/cli"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/monitor"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	processing = color.New(color.FgGreen)
	invalid    = color.New(color.Bold, color.FgRed)
)

func Run() {
	systemOs := detectOS()

	arg := os.Args[1:]
	processing.Println("➜ system-monitor starting...")

	if len(arg) == 0 {
		invalid.Println("No arguments provided.")
		os.Exit(0)
	}

	switch arg[0] {
	case "run":
		switch systemOs {
		case "linux":
			// count cpu core with logical cores
			cpuCount := monitor.CountCPUCore()

			// read cpu tasks
			startStats, err := monitor.ReadCPUTasks(cpuCount)
			if err != nil {
				invalid.Printf("Error reading CPU stats, failed to %s\n", err.Error())
				invalid.Printf("Invalid OS system, your current OS is: %s\n", systemOs)
				os.Exit(0)
			}

			// read network stats
			startNetStat, _ := monitor.ReadNetDev()

			for {
				// wait 1 second for delay
				time.Sleep(1 * time.Second)

				// read cpu tasks again after 1-second delay
				endStats, err := monitor.ReadCPUTasks(cpuCount)
				if err != nil {
					invalid.Printf("Error reading CPU stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				// read networks tasks again after 1-second delay
				endNetStat, err := monitor.ReadNetDev()
				if err != nil {
					invalid.Printf("Error reading net stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				// clear screen
				cli.ClearScreen()
				proc, err := monitor.GetProc()
				if err != nil {
					invalid.Printf("Error reading process stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}
				fmt.Printf("%+v\n", proc)
				// receive calculated network usage
				netStat, err := monitor.ReadNetUsage(startNetStat, endNetStat)
				if err != nil {
					invalid.Printf("Error reading Network stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				fmt.Printf("%+v\n", netStat)

				// read memory usage
				memResp, err := monitor.ReadMemUsage()
				if err != nil {
					invalid.Printf("Error reading Memory and Swap stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				// calculate cpu usage
				cpuUsage, err := monitor.CalculateCPUUsage(startStats, endStats)
				if err != nil {
					log.Fatalln(err)
				}

				processing.Printf("system-monitoring - %v", time.Now().Format("15:04:05"))
				fmt.Printf("%+v\n", memResp)
				// draw table
				table := cli.DrawTable(cpuUsage)

				// render table
				table.Render()

				// update startStats to endStats for the next interval calc
				startStats = endStats
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
		processing.Printf("➜ OS: %s\n", osSystem)
		processing.Printf("➜ Architecture: %s\n", osSystem)
	default:
		processing.Printf("%s.\n", osSystem)
	}

	return osSystem
}
