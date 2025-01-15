package docker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types/image"
)

type Image struct {
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	CreatedAt int64  `json:"created_at"`
	Size      int64  `json:"size"`
	ID        string `json:"id"`
}

func ListImages() ([]byte, error) {
	ctx := context.Background()
	cli := GetClient()

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
		// Handling case where RepoTags might be empty or contain only the name
		tag := ""
		if len(img.RepoTags) > 1 {
			tag = img.RepoTags[1]
		} else if len(img.RepoTags) == 1 {
			tag = "latest" // Assuming "latest" for unnamed tags
		}

		imageList = append(imageList, Image{
			Name:      img.RepoTags[0],
			Tag:       tag,
			CreatedAt: img.Created,
			Size:      img.Size,
			ID:        img.ID,
		})
	}

	// Marshal the image list to JSON
	jsonData, err := json.Marshal(imageList)
	if err != nil {
		fmt.Printf("failed to marshal to JSON: %v", err)
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return jsonData, nil
}
