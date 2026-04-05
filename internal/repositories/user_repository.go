package repositories

import (
	models "github.com/JoaoGSantiago/starti-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	ListPublicPosts(userID uint) ([]models.Post, error)
	ListPublicComments(userID uint) ([]models.Comment, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) ListPublicPosts(userID uint) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("user_id = ? AND archived = false", userID).Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// Busca comentários do usuário filtrando apenas posts não arquivados via JOIN,
// porque archived fica na tabela posts e não em comments.
func (r *userRepository) ListPublicComments(userID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.
		Joins("JOIN posts ON posts.id = comments.post_id").
		Where("comments.user_id = ? AND posts.archived = false", userID).
		Preload("User").
		Preload("Post").
		Preload("Post.User").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
