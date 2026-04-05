package services

import (
	"errors"

	"github.com/JoaoGSantiago/starti-backend/internal/model"
	"github.com/JoaoGSantiago/starti-backend/internal/repositories"
	"gorm.io/gorm"
)

type CreateCommentInput struct {
	UserID  uint   `json:"user_id" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
	Message string `json:"message" binding:"required,min=1"`
}

type UpdateCommentInput struct {
	Message string `json:"message" binding:"required,min=1"`
}

type CommentService interface {
	Create(input CreateCommentInput) (*models.Comment, error)
	Update(id uint, input UpdateCommentInput) (*models.Comment, error)
	Delete(id uint) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) Create(input CreateCommentInput) (*models.Comment, error) {
	comment := &models.Comment{
		UserID:  input.UserID,
		PostID:  input.PostID,
		Message: input.Message,
	}
	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}

	createdComment, err := s.repo.FindByID(comment.ID)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (s *commentService) Update(id uint, input UpdateCommentInput) (*models.Comment, error) {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comentario nao encontrado")
		}
		return nil, err
	}
	comment.Message = input.Message
	if err := s.repo.Update(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("comentario nao encontrado")
		}
		return err
	}
	return s.repo.Delete(id)
}
