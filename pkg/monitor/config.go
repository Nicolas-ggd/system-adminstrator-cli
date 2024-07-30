package monitor

import (
	"strings"

	. "github.com/klauspost/cpuid/v2"
)

var (
	// dir for kernel activity
	kernelActivity = "/proc/stat"
	kernelMem      = "/proc/meminfo"
	kernelNet      = "/proc/net/dev"
)

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
