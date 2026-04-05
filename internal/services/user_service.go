package services

import (
	"errors"

	models "github.com/JoaoGSantiago/starti-backend/internal/model"
	"github.com/JoaoGSantiago/starti-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Username  string `json:"username" binding:"required,min=3,max=50" example:"joaogs"`
	Name      string `json:"name" binding:"required,min=2,max=100" example:"Joao Santiago"`
	Email     string `json:"email" binding:"required,email" example:"joao@joao.com"`
	Password  string `json:"password" binding:"required,min=6" example:"senha123"`
	Biography string `json:"biography" example:"Desenvolvedor Go e APIs REST"`
}

type UpdateUserInput struct {
	Name      string `json:"name" binding:"omitempty,min=2,max=100" example:"Joao Santiago Junior"`
	Biography string `json:"biography" example:"Backend engineer focado em Go"`
}

type UserService interface {
	Create(input CreateUserInput) (*models.User, error)
	ListAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(id uint, input UpdateUserInput) (*models.User, error)
	Delete(id uint) error
	ListPublicPosts(userID uint) ([]models.Post, error)
	ListPublicComments(userID uint) ([]models.Comment, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) ListAll() ([]models.User, error) {
	return s.repo.FindAll()
}

// Verifica unicidade de email e username antes de criar,
// depois faz hash da senha com bcrypt antes de persistir.
func (s *userService) Create(input CreateUserInput) (*models.User, error) {
	if _, err := s.repo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email em uso")
	}
	if _, err := s.repo.FindByUsername(input.Username); err == nil {
		return nil, errors.New("username em uso")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  input.Username,
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashed),
		Biography: input.Biography,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario nao encontrado")
		}
		return nil, err
	}
	return user, nil
}

// Name só é atualizado se enviado; Biography sempre sobrescreve (permite limpar o campo).
func (s *userService) Update(id uint, input UpdateUserInput) (*models.User, error) {
	user, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	user.Biography = input.Biography

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *userService) ListPublicPosts(userID uint) ([]models.Post, error) {
	if _, err := s.GetByID(userID); err != nil {
		return nil, err
	}
	return s.repo.ListPublicPosts(userID)
}

func (s *userService) ListPublicComments(userID uint) ([]models.Comment, error) {
	if _, err := s.GetByID(userID); err != nil {
		return nil, err
	}
	return s.repo.ListPublicComments(userID)
}
