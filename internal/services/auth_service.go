package services

import (
	"errors"

	"github.com/JoaoGSantiago/starti-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"joao@joao.com"`
	Password string `json:"password" binding:"required" example:"senha123"`
}

type AuthService interface {
	Login(input LoginInput) (string, error)
}

type authService struct {
	userRepo   repositories.UserRepository
	jwtService JWTService
}

func NewAuthService(userRepo repositories.UserRepository, jwtService JWTService) AuthService {
	return &authService{userRepo: userRepo, jwtService: jwtService}
}

// Retorna o mesmo erro para usuário não encontrado e senha errada
// para não revelar quais emails estão cadastrados (enumeração de usuários).
func (s *authService) Login(input LoginInput) (string, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.jwtService.Generate(user.ID)
}
