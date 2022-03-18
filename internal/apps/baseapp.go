package apps

type application struct {
	name        string
	currentTag  string
	newTag      string
	buildPath   string
	composePath string
	// backupPath  string
}

type appInfo interface {
	Name() string
	CurrentTag() string
	NewTag() string
	BuildPath() string
	ComposePath() string
}

func (a *application) Name() string {
	return a.name
}

func (a *application) CurrentTag() string {
	return a.currentTag
}

func (a *application) NewTag() string {
	return a.newTag
}

func (a *application) BuildPath() string {
	return a.buildPath
}

func (a *application) ComposePath() string {
	return a.composePath
}

type applicationOption func(*application)

func WithAppName(name string) applicationOption {
	return func(a *application) {
		a.name = name
	}
}

func WithAppCurrentTag(tag string) applicationOption {
	return func(a *application) {
		a.currentTag = tag
	}
}

func WithAppNewTag(tag string) applicationOption {
	return func(a *application) {
		a.newTag = tag
	}
}

func WithAppComposePath(path string) applicationOption {
	return func(a *application) {
		a.composePath = path
	}
}

func WithAppBuildPath(path string) applicationOption {
	return func(a *application) {
		a.buildPath = path
	}
}

func NewApp(options ...applicationOption) appInfo {
	a := &application{}
	for _, option := range options {
		option(a)
	}
	return a
}

type backendApp interface {
	// application
	appInfo
}

type frontendApp interface {
	// application
	appInfo
}

type JavaApp struct {
	backendApp
}

type PythonApp struct {
	backendApp
}
type GoApp struct {
	backendApp
}
type RustApp struct {
	backendApp
}

type VersionUpdate interface {
	BuildImage()
	ModifyTag()
	Update()
}
