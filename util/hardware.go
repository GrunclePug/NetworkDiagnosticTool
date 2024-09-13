package util

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"strconv"
	"time"
)

type SysInfo struct {
	Hostname string
	Platform string
	Uptime   string
	CPU      string
	RAM      string
	Disk     string
}

func GetSysInfo() SysInfo {
	// Pull data from system
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	vmInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/") // Use '/' for Linux/Unix and '\\' for Windows

	// Format data
	uptime, _ := time.ParseDuration(strconv.FormatUint(hostInfo.Uptime, 10) + "s")
	vmUsed := fmt.Sprintf("%.2f", float32(vmInfo.Used)/1024.0/1024.0/1024.0)
	vmTotal := fmt.Sprintf("%.2f", float32(vmInfo.Total)/1024.0/1024.0/1024.0)
	vmPercent := fmt.Sprintf("%.0f", vmInfo.UsedPercent)
	diskUsed := fmt.Sprintf("%.2f", float32(diskInfo.Used)/1024.0/1024.0/1024.0)
	diskTotal := fmt.Sprintf("%.2f", float32(diskInfo.Total)/1024.0/1024.0/1024.0)
	diskPercent := fmt.Sprintf("%.0f", diskInfo.UsedPercent)

	// Build and return SysInfo Struct
	return SysInfo{
		Hostname: hostInfo.Hostname,
		Platform: hostInfo.Platform,
		Uptime:   uptime.String(),
		CPU:      cpuInfo[0].ModelName,
		RAM:      fmt.Sprintf("%v/%vGB (%v%%)", vmUsed, vmTotal, vmPercent),
		Disk:     fmt.Sprintf("%v/%vGB (%v%%)", diskUsed, diskTotal, diskPercent),
	}
}
