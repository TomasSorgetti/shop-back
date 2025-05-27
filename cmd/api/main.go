package main

import (
	"go-app/internal/infrastructure/config"
	"go-app/internal/interfaces/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    db := config.GetDB()
    defer db.Close()

    r := gin.Default()
    r.SetTrustedProxies([]string{"127.0.0.1"})
    routes.SetupRoutes(r, db)
    r.Run(":8080")
}