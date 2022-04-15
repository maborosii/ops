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

// app信息结构
#appinfo: {
	app_name: string

	package_name: string
	package_type: string | *"file" | "directory"
	package_path: string | *"\(build_path)/\(package_name)"
	if extra_build_path != _|_ {
		package_path: "\(build_path)/\(extra_build_path)/\(package_name)"
	}

	build_path: string | *"\(path.build.prod)/\(app_name)"
	if parents_app != _|_ {
		build_path: "\(path.build.prod)/\(parents_app)"
	}
	extra_build_path?: string
	parents_app?:      string

	helm_path: string | *"\(path.compose.prod)/\(app_name)"
	if parents_app != _|_ {
		helm_path: "\(path.build.prod)/\(parents_app)"
	}
	...
}
#frontendapp: #appinfo & {
	app_name:     string
	package_name: app_name
	package_type: "directory"
	parents_app:  "ale-frontend"
}

javalist: [...#appinfo]
pythonlist: [...#appinfo]
golist: [...#appinfo]
frontendlist: [...#frontendapp]

javalist: [
	#appinfo & {
		app_name:     "ale-archive-app"
		package_name: "ale-archive-app.jar"
	},
	#appinfo & {
		app_name:     "ale-case-service"
		package_name: "ale-case-service-0.0.1-SNAPSHOT.jar"
	},
	#appinfo & {
		app_name:         "ale-data-center-platform"
		package_name:     "ale-data-center-platform-0.0.1-SNAPSHOT.jar"
		extra_build_path: "ale-data-center-platform"
	},
	#appinfo & {
		app_name:     "ale-email-delivery-app"
		package_name: "ale-email-delivery-app.jar"
	},
	#appinfo & {
		app_name:     "ale-external-service"
		package_name: "ale-external-service.jar"
	},
	#appinfo & {
		app_name:     "ale-law-case-supervision"
		package_name: "ale-law-case-supervision.jar"
	},
	#appinfo & {
		app_name:     "ale-mq-consumer"
		package_name: "ale-mq-consumer-0.0.1-SNAPSHOT.jar"
	},
	#appinfo & {
		app_name:     "ale-mq-message"
		package_name: "ale-mq-message.jar"
	},
	#appinfo & {
		app_name:     "ale-mysql-data"
		package_name: "ale-mysql-data.jar"
	},
	#appinfo & {
		app_name:         "ale-office-service"
		package_name:     "ale-office-service.jar"
		extra_build_path: "ale-office-service"
	},
	#appinfo & {
		app_name:         "ale-platform"
		package_name:     "ale-platform-0.0.1-SNAPSHOT.jar"
		extra_build_path: "ale-platform"
	},
	#appinfo & {
		app_name:     "ale-seal-service"
		package_name: "ale-seal-service.jar"
	},
	#appinfo & {
		app_name:     "ale-task-job-admin"
		package_name: "ale-task-job-admin.jar"
	},
	#appinfo & {
		app_name:     "ale-task-job-executor-biz-data"
		package_name: "ale-task-job-executor-biz-data.jar"
	},
	#appinfo & {
		app_name:     "ale-task-job-executor-data-exchange"
		package_name: "ale-task-job-executor-data-exchange.jar"
	},
	#appinfo & {
		app_name:     "ale-task-job-executor-push-data-share"
		package_name: "ale-task-job-executor-push-data-share.jar"
	},
	#appinfo & {
		app_name:     "ale-task-job-executor-supervision"
		package_name: "ale-task-job-executor-supervision.jar"
	},
	#appinfo & {
		app_name:     "appr-law-card-sync-task"
		package_name: "appr-law-card-sync-task.jar"
	},
	#appinfo & {
		app_name:     "appr-law-datacenter-service-egress"
		package_name: "appr-law-egress-gateway.jar"
	},
	#appinfo & {
		app_name:     "appr-law-team-service"
		package_name: "appr-law-egress-gateway.jar"
	},
	#appinfo & {
		app_name:     "archiveweb"
		package_type: "directory"
		package_name: "ArchiveWeb"
	},
	#appinfo & {
		app_name:     "changeTool"
		package_name: "changeTool.war"
	},
	#appinfo & {
		app_name:     "docking"
		package_name: "docking.jar"
	},
	#appinfo & {
		app_name:     "esTools"
		package_type: "directory"
		package_name: "esTool"
	},
	#appinfo & {
		app_name:     "gateway"
		package_name: "appr-law-api-gateway.jar"
	},
	#appinfo & {
		app_name:     "law-app-apprtask"
		package_type: "directory"
		package_name: "tomcat-appr"
	},
	#appinfo & {
		app_name:         "law-app-attachmenttool"
		package_type:     "directory"
		package_name:     "AttachmentTool"
		extra_build_path: "webapps"
	},
	#appinfo & {
		app_name:         "law-app-cat"
		package_type:     "directory"
		package_name:     "cat"
		extra_build_path: "webapps"
	},
	#appinfo & {
		app_name:     "law-app-egress"
		package_name: "appr-law-egress-gateway.jar"
	},
	#appinfo & {
		app_name:     "xxl-job-admin"
		package_name: "xxl-job-admin.jar"
	},
]

frontendlist: [
	#frontendapp & {
		app_name: "ales"
	},
	#frontendapp & {
		app_name: "enforcement"
	},
	#frontendapp & {
		app_name: "enforcement-pad"
	},
	#frontendapp & {
		app_name: "enforcement-cloud"
	},
	#frontendapp & {
		app_name: "offsite-enforcement"
	},
	#frontendapp & {
		app_name: "offsite-ales"
	},
]

minstone: app: {
	backend: {
		java:   javalist
		python: pythonlist
	}
	frontend: frontendlist
}
