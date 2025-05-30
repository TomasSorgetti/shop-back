package routes

import (
	"database/sql"
	"go-shop/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB, container *di.Container) {
    api := r.Group("/api/v1")

	// auth routes /api/v1/auth
	SetupAuthRoutes(api, db)
}