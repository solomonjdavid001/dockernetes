package docker

import (
	"log"
	"sync"
	"time"

	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
)

var (
	once         sync.Once
	dockerClient *client.Client
)

func initDockerClient() *client.Client {
	once.Do(func() {
		var err error
		dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
	})
	return dockerClient
}

func FetchImages() ([]byte, error) {
	ctx := context.Background()
	cli := initDockerClient()

	images, err := cli.ImageList(ctx, image.ListOptions{
		All: true,
	})
	if err != nil {
		fmt.Printf("failed to list images: %v", err)
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	capacity := len(images)
	imageList := make([]Image, 0, capacity)

	for _, img := range images {
		tag := ""
		if len(img.RepoTags) > 1 {
			tag = img.RepoTags[1]
		} else if len(img.RepoTags) == 1 {
			tag = "latest"
		}

		imageList = append(imageList, Image{
			Name:      img.RepoTags[0],
			Tag:       tag,
			CreatedAt: img.Created,
			Size:      img.Size,
			ID:        img.ID,
		})
	}

	jsonData, err := json.Marshal(imageList)
	if err != nil {
		fmt.Printf("failed to marshal to JSON: %v", err)
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return jsonData, nil
}

func FetchContainers() ([]types.Container, error) {
	ctx := context.Background()
	cli := initDockerClient()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})

	if err != nil {
		fmt.Printf("failed to list containers: %v", err)
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	return containers, nil
}

func (s *SystemInfo) CalculateCpuUsage() {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("Error getting CPU usage: %s", err.Error())
		s.CPUUsage = 0
		return
	}
	s.CPUUsage = cpuUsage[0]
}

func (s *SystemInfo) CalculateRamUsage() {
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
