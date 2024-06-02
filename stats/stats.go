/**
 * 책에서는 goprocinfo를 사용하여 linux에서만 사용 가능하지만,
 * 여러 플랫폼에서 사용하기 위해 gopsutil 라이브러리로 대체
 * https://github.com/shirou/gopsutil?tab=readme-ov-file
 */
package stats

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
)

type Stats struct {
	MemStats  *mem.VirtualMemoryStat
	DiskStats *disk.UsageStat
	CpuStats  *cpu.TimesStat
	LoadStats *load.AvgStat
	TaskCount int
}

func (s *Stats) MemTotalKb() uint64 {
	return s.MemStats.Total
}

func (s *Stats) MemAvailableKb() uint64 {
	return s.MemStats.Available
}

func (s *Stats) MemUsedKb() uint64 {
	return s.MemStats.Used
}

func (s *Stats) MemUsedPercent() uint64 {
	return uint64(s.MemStats.UsedPercent)
}

func (s *Stats) DiskTotal() uint64 {
	return s.DiskStats.Total
}

func (s *Stats) DiskFree() uint64 {
	return s.DiskStats.Free
}

func (s *Stats) DiskUsed() uint64 {
	return s.DiskStats.Used
}

/*
*

	참고 : https://stackoverflow.com/questions/23367857/accurate-calculation-of-cpu-usage-given-in-percentage-in-linux
*/
func (s *Stats) CpuUsage() float64 {
	idle := s.CpuStats.Idle + s.CpuStats.Iowait
	nonIdle := s.CpuStats.User + s.CpuStats.Nice + s.CpuStats.System +
		s.CpuStats.Irq + s.CpuStats.Softirq + s.CpuStats.Steal
	total := idle + nonIdle

	if total == 0 {
		return 0.00
	}

	return (total - idle) / total
}

func GetStats() *Stats {
	return &Stats{
		MemStats:  GetMemoryInfo(),
		DiskStats: GetDiskInfo(),
		CpuStats:  GetCpuInfo(),
		LoadStats: GetLoadInfo(),
	}
}

func GetMemoryInfo() *mem.VirtualMemoryStat {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error reading memory info: %v\n", err)
		return &mem.VirtualMemoryStat{}
	}
	return memInfo
}

func GetDiskInfo() *disk.UsageStat {
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error reading disk info from /: %v\n", err)
		return &disk.UsageStat{}
	}
	return diskInfo
}

func GetCpuInfo() *cpu.TimesStat {
	cpuInfo, err := cpu.Times(false)
	if err != nil {
		log.Printf("Error reading cpu info: %v\n", err)
		return &cpu.TimesStat{}
	}
	return &cpuInfo[0]
}

func GetLoadInfo() *load.AvgStat {
	loadInfo, err := load.Avg()
	if err != nil {
		log.Printf("Error reading load info: %v\n", err)
		return &load.AvgStat{}
	}
	return loadInfo
}
