package repositories

import (
	"go-blog/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePost(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(post *models.Post) (*models.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	if err := r.db.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) DeletePost(id uint) error {
	if err := r.db.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
