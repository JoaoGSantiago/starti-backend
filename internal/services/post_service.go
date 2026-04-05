package services

import (
	"errors"

	models "github.com/JoaoGSantiago/starti-backend/internal/model"
	"github.com/JoaoGSantiago/starti-backend/internal/repositories"
	"gorm.io/gorm"
)

type CreatePostInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required,min=1"`
}

type UpdatePostInput struct {
	Text string `json:"text" binding:"required,min=1"`
}

type PostService interface {
	Create(input CreatePostInput) (*models.Post, error)
	GetByID(id uint) (*models.Post, error)
	Update(id uint, input UpdatePostInput) (*models.Post, error)
	Delete(id uint) error
	Archive(id uint) error
	ListComments(postID uint) ([]models.Comment, error)
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) Create(input CreatePostInput) (*models.Post, error) {
	post := &models.Post{
		UserID: input.UserID,
		Text:   input.Text,
	}
	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	createdPost, err := s.repo.FindByID(post.ID)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (s *postService) GetByID(id uint) (*models.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post nao encontrado")
		}
		return nil, err
	}
	return post, nil
}

func (s *postService) Update(id uint, input UpdatePostInput) (*models.Post, error) {
	post, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	post.Text = input.Text
	if err := s.repo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *postService) Archive(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Archive(id)
}

func (s *postService) ListComments(postID uint) ([]models.Comment, error) {
	if _, err := s.GetByID(postID); err != nil {
		return nil, err
	}
	return s.repo.ListComments(postID)
}
