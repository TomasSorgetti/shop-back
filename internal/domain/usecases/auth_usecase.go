// internal/domain/usecases/auth_usecase.go
package usecases

import (
	"errors"

	"go-shop/internal/domain/repository"
)

type HashService interface {
    HashPassword(password string) (string, error)
}

type JWTService interface {
    GenerateToken(userID string) (string, error)
}

type EmailService interface {
    SendWelcomeEmail(email string) error
}

type AuthUsecase interface {
    Register(email, password string) (string, error)
}

type authUsecase struct {
    userRepo     repository.UserRepository
    hashService  HashService
    jwtService   JWTService
    emailService EmailService
}

func NewAuthUsecase(userRepo repository.UserRepository, hashService HashService, jwtService JWTService, emailService EmailService) AuthUsecase {
    return &authUsecase{userRepo, hashService, jwtService, emailService}
}

func (u *authUsecase) Register(email, password string) (string, error) {

    return "", errors.New("not implemented")
}