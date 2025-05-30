package repository

import "go-shop/internal/domain/entities"

type UserRepository interface {
    Create(user *entities.User) error
    ExistsByEmail(email string) (bool, error)
}

