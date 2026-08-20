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
	
	adapters.GetVersion()
	server.ServerStart()


	// req := models.DownloadRequest{
	// 	URL:        "https://youtu.be/...",
	// 	Type:       "video",
	// 	Resolution: "1070",
	// }

	args := core.BuildArgs(json_tests.Json_test())
	processedArgs := fmt.Sprintf("%+v", args)

	fmt.Printf("This is the way to go: %s", processedArgs)
	fmt.Println("This part shouldn't be reached")
	

}
