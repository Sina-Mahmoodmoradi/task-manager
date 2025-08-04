package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    log.Println("🚀 Starting server on http://localhost:8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
