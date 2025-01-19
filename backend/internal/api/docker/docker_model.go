package docker

type Image struct {
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	CreatedAt int64  `json:"created_at"`
	Size      int64  `json:"size"`
	ID        string `json:"id"`
}

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
