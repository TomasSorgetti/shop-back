package persistance

import (
	"database/sql"
	"go-shop/internal/infrastructure/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
    DB *sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
    db, err := sql.Open("mysql", cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to MySQL: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping MySQL: %v", err)
    }

    return &Database{DB: db}, nil
}

func (d *Database) Close() error {
    return d.DB.Close()
}