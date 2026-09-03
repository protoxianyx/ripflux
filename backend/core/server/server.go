package server

import (
	"net/http"
	"strings"

	// "go/format"
	"ripflux/config"
	"ripflux/core"
	"ripflux/models"

	// "ripflux/core/server"

	"ripflux/core/adapters"
	// "ripflux/utils/adapters/ytdlpexe"
	"ripflux/utils/loggers"
	packagemanager "ripflux/utils/package"

	// "ripflux/utils/adapters"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func downloadHandler(c *gin.Context) {
	var req models.DownloadRequestModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if err := ProcessDownloadRequest(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Download failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Sucess",
		"message": "Download Complete sucessfully",
	})

}

func versionHandler(c *gin.Context) {
	latestVersion, err := packagemanager.GetLatestVersionInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	version, err := adapters.GetVersion()
	if err != nil {
		loggers.Logf("GetVersion failed: %v\n", err)
		version = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// install := packagemanager.Install()
	loggers.Logf("\nInstalled Version: %v\nLatest Version: %v\n", version, latestVersion)

	c.JSON(200, gin.H{
		"version":       version,
		"latestVersion": latestVersion,
	})
}

func updaterHandler(c *gin.Context) {
	packagemanager.Install()
	loggers.TaskLog(config.TEST_LOG_FILE_PATH, "Successfully reaching the updatehandler")

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update completed successfully",
	})
}

func latestVersionInfo(c *gin.Context) {
	latestVersionInfo, err := packagemanager.GetLatestVersionInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"latestVersionInfo": latestVersionInfo,
	})

}

func ServerStart() {

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/download", downloadHandler)
	router.GET("/version", versionHandler)
	router.POST("/latestVersion", updaterHandler)
	router.GET("/latestVersionInfo", latestVersionInfo)

	router.Run(":8080")
}

func ProcessDownloadRequest(submittedRequestData models.DownloadRequestModel) error {

	submittedRequestData.Format = strings.ToLower(submittedRequestData.Format)
	submittedRequestData.Resolution = strings.ToLower(submittedRequestData.Resolution)

	args := core.BuildDownloadCommand(submittedRequestData)

	loggers.TaskLog(config.INPUT_LOG_FILE_PATH, args)

	return adapters.Download(args)

	// return fmt.Errorf("downloader was skipped: adapters.Download is disabled")
}

func newFunc() {

}
