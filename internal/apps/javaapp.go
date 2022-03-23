package apps

type JavaApp struct {
	*BackendApp

	jarName string
}

func (j *JavaApp) SetJarName(name string) {
	j.jarName = name
}

func (j *JavaApp) GetName() string {
	return j.jarName
}

func NewJavaApp(jarName string) AppInfo {
	return &JavaApp{
		BackendApp: &BackendApp{
			&BaseApp{},
		},
		jarName: jarName,
	}
}
