package system

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type DiskInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type SystemInfo struct {
	CPUUsage    float64    `json:"cpu_usage"`
	MemoryUsage MemoryInfo `json:"memory_usage"`
	DiskUsage   DiskInfo   `json:"disk_usage"`
	HostName    string     `json:"hostname"`
}

func (s *SystemInfo) cpuUsage() {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("Error getting CPU usage: %s", err.Error())
		s.CPUUsage = 0
		return
	}
	s.CPUUsage = cpuUsage[0]
}

func (s *SystemInfo) ramUsage() {
	memUsage, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory usage: %s", err.Error())
		s.MemoryUsage = MemoryInfo{}
	}

	s.MemoryUsage = MemoryInfo{
		Total:       memUsage.Total / 1024 / 1024 / 1024,
		Used:        memUsage.Used / 1024 / 1024 / 1024,
		UsedPercent: memUsage.UsedPercent,
	}
}

func GetSystemInfo(c *gin.Context) {
	systemInfo := SystemInfo{}

	// Get CPU Usage
	systemInfo.cpuUsage()

	// Get RAM Usage
	systemInfo.ramUsage()

	c.JSON(http.StatusOK, gin.H{
		"systemInfo": systemInfo,
	})
}
