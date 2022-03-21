package apps

type JavaApp struct {
	*BackendApp

	jarName string
}

func (j *JavaApp) SetJar(name string) {
	j.jarName = name

}
