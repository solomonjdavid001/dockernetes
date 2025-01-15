package docker

import (
	"sync"

	"github.com/docker/docker/client"
)

var (
	once         sync.Once
	dockerClient *client.Client
)

func GetClient() *client.Client {
	once.Do(func() {
		var err error
		dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
	})
	return dockerClient
}
