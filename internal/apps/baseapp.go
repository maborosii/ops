package apps

type BaseApp struct {
	// tag         string
	name        string
	project     string
	currentTag  string
	newTag      string
	buildPath   string
	composePath string
	extraPath   string
	// backupPath  string
}

// type AppInfo interface {
// 	SetName(name string)
// 	SetCurrentTag(tag string)
// 	SetNewTag(tag string)
// 	SetBuildPath(path string)
// 	SetComposePath(path string)
// }

// func (b *BaseApp) SetName(name string) {
// 	b.name = name
// }

// func (b *BaseApp) SetCurrentTag(tag string) {
// 	b.currentTag = tag
// }

// func (b *BaseApp) SetNewTag(tag string) {
// 	b.newTag = tag
// }

// func (b *BaseApp) SetBuildPath(path string) {
// 	b.buildPath = path
// }

// func (b *BaseApp) SetComposePath(path string) {
// 	b.composePath = path
// }
type AppInfo interface {
	GetName() string
	GetRealName() string
	GetProject() string
	GetCurrentTag() string
	GetNewTag() string
	GetBuildPath() string
	GetComposePath() string
	GetExtraPath() string
	// SetTag(tag string)
	// GetTag() string
}

func (b *BaseApp) GetName() string {
	return b.name
}

func (b *BaseApp) GetRealName() string {
	return b.name
}

func (b *BaseApp) GetProject() string {
	return b.project
}

func (b *BaseApp) GetCurrentTag() string {
	return b.currentTag
}

func (b *BaseApp) GetNewTag() string {
	return b.newTag
}

func (b *BaseApp) GetBuildPath() string {
	return b.buildPath
}

func (b *BaseApp) GetComposePath() string {
	return b.composePath
}
func (b *BaseApp) GetExtraPath() string {
	return b.extraPath
}

// func (b *BaseApp) SetTag(tag string) {
// 	b.tag = tag
// }
// func (b *BaseApp) GetTag() string {
// 	return b.tag
// }
