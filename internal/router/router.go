package router

import (
    "github.com/gin-gonic/gin"
    "github.com/Sina-Mahmoodmoradi/task-manager/internal/handler"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", handler.PingHandler)

    return r
}
