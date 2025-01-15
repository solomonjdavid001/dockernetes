package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func GetContainers() ([]types.Container, error) {
	ctx := context.Background()
	cli := GetClient()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})

	if err != nil {
		fmt.Printf("failed to list containers: %v", err)
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	return containers, nil
}
