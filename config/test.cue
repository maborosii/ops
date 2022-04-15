package config

// 公共变量
path: {
	build: {
		prod: "/data/yw_opert/docker/build/app"
	}
	compose: {
		prod: "/data/yw_opert/k8s/prod/helm_prod"
		qa:   "/data/yw_opert/k8s/qa/helm_qa"
	}
}

#appinfo: {
	app:                 string
	"type":              string | *"file" | "directory"
	"package_name":      string
	"build_path":        string | *"\(path.build.prod)/\(app)"
	"extra_build_path"?: string
	"parents_app"?:      string
	"helm_path":         string | *"\(path.compose.prod)/\(app)"
	...
}

test: #appinfo & {
	app:            "test"
	"package_name": "test.jar"
}
