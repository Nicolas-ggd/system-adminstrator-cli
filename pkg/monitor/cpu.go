package monitor

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/klauspost/cpuid/v2"
)

var (
	processing = color.New(color.Bold, color.FgGreen)
)

func GetLinuxCPU() (idle, total uint64) {
	contents, err := os.ReadFile(kernelDir)
	if err != nil {
		log.Fatalf("failed to read cpu stats: %v", err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}

	return
}

func CpuLogger() {
	processing.Println("➜ Name:", CPU.BrandName)
	processing.Println("➜ PhysicalCores:", CPU.PhysicalCores)
	processing.Println("➜ ThreadsPerCore:", CPU.ThreadsPerCore)
	processing.Println("➜ LogicalCores:", CPU.LogicalCores)
	processing.Println("➜ Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	processing.Println("➜ Features:", strings.Join(CPU.FeatureSet(), ","))
	processing.Println("➜ CacheLine bytes:", CPU.CacheLine)
	processing.Println("➜ L1 Data Cache:", CPU.Cache.L1D, "bytes")
	processing.Println("➜ L1 Instruction Cache:", CPU.Cache.L1I, "bytes")
	processing.Println("➜ L2 Cache:", CPU.Cache.L2, "bytes")
	processing.Println("➜ L3 Cache:", CPU.Cache.L3, "bytes")
	processing.Println("➜ Frequency", CPU.Hz, "hz")
}
