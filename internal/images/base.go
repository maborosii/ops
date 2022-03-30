package images

import (
	"context"
	"ops/internal/apps"
)

type ImagesItf interface {
	SetInChan(inchan <-chan apps.AppInfo)
	SetOutChan(outchan chan<- BuildInfo)
	SetImageTags(tags ...string)
	GetImageTags() []string
	SetRepository(rep string)
	GetRepository() string

	BuildImage(context.Context, apps.AppInfo) error
	PushImage(context.Context, apps.AppInfo) error
	// CleanImage() error
	// ModifyBuildFile() error

	Run(context.Context)
}

type BuildInfo struct {
	apps.AppInfo
	ImagesItf
}

func NewBuildInfo(a apps.AppInfo, images ImagesItf) BuildInfo {
	return BuildInfo{a, images}

}

type BaseImage struct {
	inChan  <-chan apps.AppInfo
	outChan chan<- BuildInfo

	tags       []string
	repository string
}

func (b *BaseImage) SetInChan(inchan <-chan apps.AppInfo) {
	b.inChan = inchan
}

func (b *BaseImage) SetOutChan(outchan chan<- BuildInfo) {
	b.outChan = outchan
}

func (b *BaseImage) SetImageTags(tags ...string) {
	b.tags = tags
}

func (b *BaseImage) GetImageTags() []string {
	return b.tags
}

func (b *BaseImage) SetRepository(rep string) {
	b.repository = rep
}

func (b *BaseImage) GetRepository() string {
	return b.repository
}
