package nux

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func TestInfo(t *testing.T) {
	procs, _ := process.Processes()
	fmt.Println(procs)

	for _, proc := range procs {
		pName, _ := proc.Name()
		if strings.Contains(pName, "go") {
			fmt.Println(proc.Name())
			fmt.Println(proc.CPUPercent())
			fmt.Println(proc.MemoryPercent())
		}
	}
	// proc, _ := process.NewProcess(20292)
	// fmt.Println(proc.Name())
	// fmt.Println(proc.CPUPercent())
	// fmt.Println(proc.MemoryPercent())
}
