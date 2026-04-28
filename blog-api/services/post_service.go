package services

import (
	"go-blog/models"
	"go-blog/repositories"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts(page, limit int) ([]models.Post, int64, error)
	UpdatePost(id uint, post *models.Post) (*models.Post, error)
	DeletePost(id uint) error
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	return s.repo.CreatePost(post)
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.GetPostByID(id)
}

func (s *postService) GetAllPosts(page, limit int) ([]models.Post, int64, error) {
	offset := (page - 1) * limit
	return s.repo.GetAllPosts(limit, offset)
}

func (s *postService) UpdatePost(id uint, post *models.Post) (*models.Post, error) {
	existingPost, err := s.repo.GetPostByID(id)

	if err != nil {
		return nil, err
	}

	existingPost.Title = post.Title
	existingPost.Content = post.Content

	return s.repo.UpdatePost(existingPost)
}

func (s *postService) DeletePost(id uint) error {
	_, err := s.repo.GetPostByID(id)

	if err != nil {
		return err
	}

	return s.repo.DeletePost(id)
}
