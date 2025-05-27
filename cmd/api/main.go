package main

import (
	"go-shop/internal/infrastructure/config"
	"go-shop/internal/infrastructure/persistance"
	"go-shop/internal/interfaces/http/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    db, err := persistance.NewDatabase(cfg)
    if err != nil {
        log.Fatalf("Error connecting to MySQL: %v", err)
    }
    defer db.Close()
    
    r := gin.Default()
    r.SetTrustedProxies([]string{"127.0.0.1"})

    routes.SetupRoutes(r)
    
    r.Run(":8080")
}