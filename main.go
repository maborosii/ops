package main

import (
	"fmt"
	"log"
	"os"

	"ops/internal/deployers/customhelm"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	// ctx := context.Background()
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }
	// path := "/home/bonbon/docker_apollo/build"
	// // buildFile, _ := os.Open(path)
	// // defer buildFile.Close()

	// buildCtx, _ := archive.TarWithOptions(path, &archive.TarOptions{})

	// s, err := cli.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
	// 	Tags: []string{"harbor.minstone.com:5002/app/node-metrics-go:dev-1"},
	// 	// Dockerfile: "Dockerfile",
	// })
	// fmt.Println(err)

	// io.Copy(os.Stdout, s.Body)

	// //
	// //

	// ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	// var authConfig = types.AuthConfig{
	// 	Username:      "admin",
	// 	Password:      "minstone123",
	// 	ServerAddress: "harbor.minstone.com:5002",
	// }

	// authConfigBytes, _ := json.Marshal(authConfig)
	// authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	// tag := "harbor.minstone.com:5002/app/node-metrics-go:dev-1"
	// opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	// rd, err := cli.ImagePush(ctxTimeout, tag, opts)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer rd.Close()
	settings := cli.New()

	actionConfig := new(action.Configuration)
	// // You can pass an empty string instead of settings.Namespace() to list
	// // all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), "qa", os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	// client := action.NewList(actionConfig)
	// // Only list deployed
	// client.Deployed = true
	// results, err := client.Run()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	os.Exit(1)
	// }

	// for _, rel := range results {
	// 	log.Printf("%+v", rel)
	// }
	client := action.NewInstall(actionConfig)
	release, err := customhelm.RunInstall([]string{"ale-case-service", "/data/yw_opert/k8s/qa/helm_qa/ale-case-service"},
		client,
		map[string]interface{}{
			"image.repository": "harbor.minstone.com/app/ale-case-service",
			"image.tag":        "1.2.4.2",
		},
		"qa",
		os.Stdout)

	if err != nil {
		fmt.Println(err)
		panic("install failed")
	}
	fmt.Println(release.Name, "install successful")

}
