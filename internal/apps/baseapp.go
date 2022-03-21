package apps

type BaseApp struct {
	name        string
	currentTag  string
	newTag      string
	buildPath   string
	composePath string
	// backupPath  string
}

type AppInfo interface {
	SetName(name string)
	SetCurrentTag(tag string)
	SetNewTag(tag string)
	SetBuildPath(path string)
	SetComposePath(path string)
}

func (b *BaseApp) SetName(name string) {
	b.name = name
}

func (b *BaseApp) SetCurrentTag(tag string) {
	b.currentTag = tag
}

func (b *BaseApp) SetNewTag(tag string) {
	b.newTag = tag
}

func (b *BaseApp) SetBuildPath(path string) {
	b.buildPath = path
}

func (b *BaseApp) SetComposePath(path string) {
	b.composePath = path
}
