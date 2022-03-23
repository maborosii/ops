package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	path := "/home/bonbon/docker_apollo/build"
	// buildFile, _ := os.Open(path)
	// defer buildFile.Close()

	buildCtx, _ := archive.TarWithOptions(path, &archive.TarOptions{})

	s, _ := cli.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		// Dockerfile: "Dockerfile",
	})

	io.Copy(os.Stdout, s.Body)
	// images, err := cli.ImageList(ctx, types.ImageListOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// cli.ImagePush()

	// for _, image := range images {
	// 	fmt.Println(image.ID, image.RepoTags)
	// }
}
