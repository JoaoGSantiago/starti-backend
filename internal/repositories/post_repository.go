package repositories

import (
	"github.com/JoaoGSantiago/starti-backend/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindByID(id uint) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
	Archive(id uint) error
	ListComments(postID uint) ([]models.Comment, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.Preload("User").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *postRepository) Archive(id uint) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Update("archived", true).Error
}

func (r *postRepository) ListComments(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("post_id = ?", postID).Preload("User").Preload("Post").Preload("Post.User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
