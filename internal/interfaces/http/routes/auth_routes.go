package routes

import (
	"database/sql"
	"go-shop/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(api *gin.RouterGroup, db *sql.DB ) {
    authHandler := handlers.NewAuthHandler()

    auth := api.Group("/auth")
    {
		// @POST /api/v1/auth/signin
        auth.POST("/signin", authHandler.Login)
		// @POST /api/v1/auth/signup
        auth.POST("/signup", authHandler.Register)
    }
}