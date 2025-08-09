package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/handler"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/service"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/infrastructure/database"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/infrastructure/security"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/middleware"
)


func SetupRouter(db *database.Database) *gin.Engine {
    r := gin.Default()
    
	// Repositories
	userRepo := repository.NewUserRepository(db.DB)
	taskRepo := repository.NewTaskRepository(db.DB)

	// Token Manager
	tokenManager := security.NewJWTTokenManager()
	
    // Services
	userService := service.NewUserService(userRepo, tokenManager)
	taskService := service.NewTaskService(taskRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	taskHandler := handler.NewTaskHandler(taskService)

	// Routes
	api := r.Group("/api/v1")
	{
		api.GET("/ping", handler.PingHandler)
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)

		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware(tokenManager))
		{
			auth.GET("/me", userHandler.GetCurrentUser)
			auth.GET("/tasks", taskHandler.ListTasks)
			auth.GET("/tasks/:id", taskHandler.GetTask)
			auth.PUT("/tasks/:id", taskHandler.UpdateTask)
			auth.POST("/tasks", taskHandler.CreateTask)
			auth.DELETE("/tasks/:id", taskHandler.DeleteTask)
		}
	}

	return r
}
