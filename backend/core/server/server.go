package server

import (
	"fmt"
	// "go/format"
	"ripflux/core"
	"ripflux/core/models"

	// "ripflux/core/server"

	"ripflux/utils/adapters"
	"ripflux/utils/loggers"

	// "ripflux/utils/adapters"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func downloadHandler(c *gin.Context) {
	var req models.DownloadRequestModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fmt.Printf("%+v\n", req)

	InputDataSend(req)

	c.JSON(200, gin.H{
		"status": "ok",
	})

}

func versionHandler(c *gin.Context) {
    version, err := adapters.GetVersion()
    if err != nil {
        loggers.Logf("GetVersion failed: %v\n", err)
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "version": version,
    })
}

func ServerStart() {

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/download", downloadHandler)
    router.GET("/version", versionHandler)

	router.Run(":8080")
}

func InputDataSend(submittedRequestData models.DownloadRequestModel) {

	VidoeFormat := submittedRequestData.URL
	fmt.Printf("This is the URL: %s\n", VidoeFormat)
	loggers.Logf("This is the URL: %s", VidoeFormat)

	args := core.BuildArgs(submittedRequestData)

	loggers.Log(args)
	loggers.Logf("Number of args: %d\n", len(args))
	for i, a := range args {
		loggers.Logf("args[%d]: %s\n", i, a)
	}

	adapters.Download(args)

}
