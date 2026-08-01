package main

import (
	// "encoding/json"
	"fmt"
	"ripflux/core"
	"ripflux/core/server"
	"ripflux/tests"

	// "ripflux/core/models"
	"ripflux/utils/adapters"
)

// "log"

// "ripflux/utils/coreinstaller"

func main() {
	// coreinstaller.Install()

	server.ServerStart()

	adapters.GetVersion()

	// req := models.DownloadRequest{
	// 	URL:        "https://youtu.be/...",
	// 	Type:       "video",
	// 	Resolution: "1070",
	// }

	args := core.BuildArgs(json_tests.Json_test())

	fmt.Println(args)



}
