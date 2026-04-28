package services

import (
	"errors"
	"fmt"
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
	postRepo repositories.PostRepository
	userRepo repositories.UserRepository
}

//func NewPostService(repo repositories.PostRepository) PostService {
//	return &postService{postRepo: repo}
//}

func NewPostService(postRepo repositories.PostRepository, userRepo repositories.UserRepository) PostService {
	return &postService{
		postRepo: postRepo,
		userRepo: userRepo, // Assign the dependency here
	}
}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	userExists, err := s.userRepo.ExistsById(post.AuthorID)

	fmt.Println(userExists)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, errors.New("author not found")
	}

	return s.postRepo.CreatePost(post)
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.postRepo.GetPostByID(id)
}

func (s *postService) GetAllPosts(page, limit int) ([]models.Post, int64, error) {
	offset := (page - 1) * limit
	return s.postRepo.GetAllPosts(limit, offset)
}

func (s *postService) UpdatePost(id uint, post *models.Post) (*models.Post, error) {
	existingPost, err := s.postRepo.GetPostByID(id)

	if err != nil {
		return nil, err
	}

	existingPost.Title = post.Title
	existingPost.Content = post.Content

	return s.postRepo.UpdatePost(existingPost)
}

func (s *postService) DeletePost(id uint) error {
	_, err := s.postRepo.GetPostByID(id)

	if err != nil {
		return err
	}

	return s.postRepo.DeletePost(id)
}
