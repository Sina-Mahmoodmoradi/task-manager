package main

import (
    "log"

    "github.com/Sina-Mahmoodmoradi/task-manager/internal/router"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/infrastructure/database"
)

func main() {
    db, err := database.NewDatabase()
    if err != nil {
        log.Fatalf("could not initialize database: %v", err)
    }
    r := router.SetupRouter(db)

    log.Println("🚀 Server running at http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
