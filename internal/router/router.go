package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/handler"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/service"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository"
	"github.com/Sina-Mahmoodmoradi/task-manager/pkg/database"
)


func SetupRouter(db *database.Database) *gin.Engine {
    r := gin.Default()
    
	// Repositories
	userRepo := repository.NewUserRepository(db.DB)

	// Services
	userService := service.NewUserService(userRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
    
	// Routes
	api := r.Group("/api/v1")
	{
		api.GET("/ping", handler.PingHandler)
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	return r
}
