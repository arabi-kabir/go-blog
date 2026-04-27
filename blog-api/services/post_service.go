package services

import (
	"go-blog/models"
	"go-blog/repositories"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(id uint, post *models.Post) (*models.Post, error)
	DeletePost(id uint) error
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo: repo}
}

// PostService methods here...
func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	return s.repo.CreatePost(post)
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.GetPostByID(id)
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
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
	return s.repo.DeletePost(id)
}
