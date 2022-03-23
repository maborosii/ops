package images

import "ops/internal/apps"

type ImagesItf interface {
	BuildImage() error
	PushImage() error
	ModifyBuildFile() error
}

type BaseImage struct {
	inChan  <-chan apps.AppInfo
	outChan chan<- apps.AppInfo
}
