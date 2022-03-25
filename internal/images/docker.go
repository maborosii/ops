package images

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"ops/internal/apps"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type DockerCfg struct {
	Repository string
	AuthCfg    types.AuthConfig

	Cli *client.Client
}

type DockerImage struct {
	*BaseImage
	*DockerCfg
}

func (di *DockerImage) BuildImage(ctx context.Context, a apps.AppInfo) error {
	buildPath := a.GetBuildPath()

	// build directory switch to tar
	buildTar, err := archive.TarWithOptions(buildPath, &archive.TarOptions{})
	if err != nil {
		return err
	}
	s, err := di.Cli.ImageBuild(ctx, buildTar, types.ImageBuildOptions{
		Tags: di.tags,
	})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, s.Body)
	return nil
}

func (di *DockerImage) PushImage(ctx context.Context, a apps.AppInfo) error {

	ctxTimeOut, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	authConfigBytes, err := json.Marshal(di.AuthCfg)
	if err != nil {
		return nil
	}

	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	tag := fmt.Sprintf("%s/%s/%s:%s", di.Repository, a.GetProject(), a.GetName(), di.tags[0])
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	rd, err := di.Cli.ImagePush(ctxTimeOut, tag, opts)

	io.Copy(os.Stdout, rd)
	if err != nil {
		return err
	}
	return nil
}

func (di *DockerImage) Run(ctx context.Context) {
	var (
		ok      bool
		appInfo apps.AppInfo
		err     error
	)

HANDLE_LOOP:
	for {
		select {
		case <-ctx.Done():
			break HANDLE_LOOP
		case appInfo, ok = <-di.inChan:
			if !ok {
				break HANDLE_LOOP
			}
		}
		// main logic
		// TODO：buildimage 和 pushimage 也需要解耦
		err = di.BuildImage(ctx, appInfo)
		if err != nil {
			panic(err)
		}
		err = di.PushImage(ctx, appInfo)
		if err != nil {
			panic(err)
		}

		di.outChan <- appInfo
	}

}
