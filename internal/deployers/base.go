package deployers

import (
	"context"
	"ops/internal/images"
)

type DeployersItf interface {
	SetInChan(inchan <-chan images.BuildInfo)
	SetOutChan(outchan chan<- images.BuildInfo)
	// SetDeployTags(tag string)

	Install(context.Context, images.BuildInfo) error
	Uninstall(images.BuildInfo) error

	Run(context.Context)
}

type BaseDeployer struct {
	inChan  <-chan images.BuildInfo
	outChan chan<- images.BuildInfo
}

func (b *BaseDeployer) SetInChan(inchan <-chan images.BuildInfo) {
	b.inChan = inchan
}

func (b *BaseDeployer) SetOutChan(outchan chan<- images.BuildInfo) {
	b.outChan = outchan
}
