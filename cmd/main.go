package main

import (
    "log"

    "github.com/Sina-Mahmoodmoradi/task-manager/internal/router"
)

func main() {
    r := router.InitRouter()

    log.Println("🚀 Server running at http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
