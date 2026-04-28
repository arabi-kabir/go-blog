package repositories

import (
	"go-blog/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts(limit, offset int) ([]models.Post, int64, error)
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
	result := r.db.Create(post)

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (r *postRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post

	result := r.db.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).
		Preload("Category").
		First(&post, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (r *postRepository) GetAllPosts(limit, offset int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	totalResult := r.db.Model(&models.Post{}).Count(&total)
	if totalResult.Error != nil {
		return nil, 0, totalResult.Error
	}

	// Get paginated data with relations
	result := r.db.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).
		Preload("Category").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return posts, total, nil
}

func (r *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	result := r.db.Save(post)

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (r *postRepository) DeletePost(id uint) error {
	result := r.db.Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
