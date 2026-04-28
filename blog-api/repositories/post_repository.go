package repositories

import (
	"context"
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
	ctx := context.Background()
	postDB := gorm.G[models.Post](r.db)

	err := postDB.Create(ctx, post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) GetPostByID(id uint) (*models.Post, error) {
	ctx := context.Background()
	postDB := gorm.G[models.Post](r.db)

	post, err := postDB.
		Preload("Author", func(db gorm.PreloadBuilder) error {
			db.Select("id", "username", "email")
			return nil
		}).
		Preload("Category", nil).
		Where("id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) GetAllPosts(limit, offset int) ([]models.Post, int64, error) {
	ctx := context.Background()
	postDB := gorm.G[models.Post](r.db)

	var total int64
	totalResult := r.db.Model(&models.Post{}).Count(&total)
	if totalResult.Error != nil {
		return nil, 0, totalResult.Error
	}

	posts, err := postDB.
		Preload("Category", nil).
		Preload("Author", nil).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(ctx)

	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	ctx := context.Background()
	postDB := gorm.G[models.Post](r.db)

	_, err := postDB.Where("id = ?", post.ID).Updates(ctx, models.Post{
		Title:       post.Title,
		Content:     post.Content,
		IsPublished: post.IsPublished,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) DeletePost(id uint) error {
	ctx := context.Background()
	postDB := gorm.G[models.Post](r.db)

	_, err := postDB.Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
