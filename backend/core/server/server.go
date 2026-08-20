package server

import (
	"fmt"
	// "go/format"
	"ripflux/core/models"
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

func ServerStart() {

    router := gin.Default()

    router.Use(cors.Default())

    router.POST("/download", downloadHandler)

    router.Run(":8080")
}

func InputDataSend(submittedRequestData models.DownloadRequestModel)  {
    

    VidoeFormat := submittedRequestData.URL
    fmt.Printf("This is the format %s\n", VidoeFormat)

    // return VidoeFormat
}

