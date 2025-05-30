package di

import (
	"database/sql"
	"go-shop/internal/interfaces/http/handlers"
)

type Container struct {
    AuthHandler *handlers.AuthHandler
    UserHandler *handlers.UserHandler
    // Otros handlers o servicios
}

func NewContainer(db *sql.DB) *Container {
    // Handlers
    authHandler := handlers.NewAuthHandler()

    return &Container{
        AuthHandler: authHandler,
    }
}