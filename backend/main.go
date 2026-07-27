package main

import (
	"fmt"
	"ripflux/core"
	"ripflux/core/models"
	"ripflux/utils/adapters"
)

// "log"

// "ripflux/utils/coreinstaller"

func main() {
	// coreinstaller.Install()

	adapters.GetVersion()

	req := models.DownloadRequest{
		URL:        "https://youtu.be/...",
		Type:       "video",
		Resolution: "1070",
	}

    args:= core.BuildArgs(req)

    fmt.Println(args)

}
