package setting

type BaseAppCfg struct {
	appName        string `json:"app_name"`
	packageName    string `json:"package_name"`
	packageType    string `json:"package_type"`
	packagePath    string `json:"package_path"`
	buildPath      string `json:"build_path"`
	extraBuildPath string `json:"extra_build_path,omitempty"`
	parentsApp     string `json:"parents_app,omitempty"`
	helmPath       string `json:"helm_path"`
}
