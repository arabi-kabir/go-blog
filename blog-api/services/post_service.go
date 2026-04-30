package services

import (
	"errors"
	"fmt"
	"go-blog/models"
	"go-blog/pkg/cache"
	"go-blog/repositories"
	"time"
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
	cache    *cache.Cache
}

func NewPostService(postRepo repositories.PostRepository, userRepo repositories.UserRepository, cache *cache.Cache) PostService {
	return &postService{
		postRepo: postRepo,
		userRepo: userRepo,
		cache:    cache,
	}
}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	userExists, err := s.userRepo.ExistsById(post.AuthorID)

	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, errors.New("author not found")
	}

	createdPost, err := s.postRepo.CreatePost(post)
	if err == nil {
		s.cache.InvalidateTag("posts")
	}

	return createdPost, err
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	cacheKey := fmt.Sprintf("post:%d", id)

	var post models.Post

	// 1. try cache
	found, _ := s.cache.Get(cacheKey, &post)
	if found {
		return &post, nil
	}

	// 2. db
	postPtr, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	// 3. cache with tags
	s.cache.Set(cacheKey, postPtr, 10*time.Minute, []string{
		"posts",
		fmt.Sprintf("post:%d", id),
	})

	return postPtr, nil
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
	existingPost.PostCategoryID = post.PostCategoryID
	existingPost.IsPublished = post.IsPublished

	return s.postRepo.UpdatePost(existingPost)
}

func (s *postService) DeletePost(id uint) error {
	_, err := s.postRepo.GetPostByID(id)

	if err != nil {
		return err
	}

	return s.postRepo.DeletePost(id)
}
