package main

import (
    "log"

    "github.com/Sina-Mahmoodmoradi/task-manager/internal/router"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/config"
)

func main() {
    config.InitDB()
    r := router.InitRouter()

    log.Println("🚀 Server running at http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
