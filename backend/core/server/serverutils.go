package server

import (
	"net/http"
	"path/filepath"
	"ripflux/config/paths"
	"ripflux/utils"

	"github.com/gin-gonic/gin"
)

func clearTmpHandler(c *gin.Context) {
	tmpDir := filepath.Join(paths.PATH_GO_BACK, paths.TMP_FOLDER)
	if err := utils.CleanTmp(tmpDir); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Temporary files cleared"})
}
