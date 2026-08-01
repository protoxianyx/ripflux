package server

import (
	"fmt"
	"ripflux/core/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func downloadHandler(c *gin.Context) {
    var req models.DownloadRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

	fmt.Printf("%+v\n", req)

    c.JSON(200, gin.H{
        "status": "ok",
    })
}

func ServerStart() {

    router := gin.Default()

    router.Use(cors.Default())

    router.POST("/download", downloadHandler)

    router.Run(":8080")
}