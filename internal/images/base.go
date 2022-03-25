package images

import (
	"context"
	"ops/internal/apps"
)

type ImagesItf interface {
	SetInChan(inchan <-chan apps.AppInfo)
	SetOutChan(outchan chan<- apps.AppInfo)
	SetImageTags(tags ...string)

	BuildImage(context.Context, apps.AppInfo) error
	PushImage(context.Context, apps.AppInfo) error
	// CleanImage() error
	// ModifyBuildFile() error

	Run(apps.AppInfo)
}

type BaseImage struct {
	inChan  <-chan apps.AppInfo
	outChan chan<- apps.AppInfo

	tags []string
}

func (b *BaseImage) SetInChan(inchan <-chan apps.AppInfo) {
	b.inChan = inchan
}

func (b *BaseImage) SetOutChan(outchan chan<- apps.AppInfo) {
	b.outChan = outchan
}

func (b *BaseImage) SetImageTags(tags ...string) {
	b.tags = tags
}
