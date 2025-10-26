package main

import (
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

type Input struct {
    Text string `json:"text"`
}

func main() {
    r := gin.Default()

    // Enable CORS (for frontend on localhost:3000)
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{"POST", "GET"},
        AllowHeaders: []string{"Content-Type"},
    }))

    // Your first route
    r.POST("/api/send", func(c *gin.Context) {
        var input Input
        if err := c.BindJSON(&input); err != nil {
            c.JSON(400, gin.H{"error": "invalid JSON"})
            return
        }

        fmt.Println("Received from frontend:", input.Text)
        c.JSON(200, gin.H{"message": "Got it!", "received": input.Text})
    })

    r.Run(":8080")
}
