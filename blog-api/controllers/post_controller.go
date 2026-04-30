package controllers

import (
	"go-blog/dto"
	"go-blog/models"
	"go-blog/pkg/response"
	"go-blog/services"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{service: service}
}

func (ctrl *PostController) CreatePost(c echo.Context) error {
	req := new(dto.CreatePostRequest)

	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	post := models.Post{
		Title:          req.Title,
		Content:        req.Content,
		IsPublished:    req.IsPublished,
		AuthorID:       req.AuthorID,
		PostCategoryID: req.PostCategoryID,
	}

	createdPost, err := ctrl.service.CreatePost(&post)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to create post", err.Error())
	}

	res := dto.PostResponse{
		ID:           createdPost.ID,
		Title:        createdPost.Title,
		Content:      createdPost.Content,
		IsPublished:  createdPost.IsPublished,
		Author:       createdPost.Author.Username,
		PostCategory: createdPost.Category.Title,
	}

	return response.Success(c, "Post created successfully", res, 201)
}

func (ctrl *PostController) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid post ID", err.Error())
	}

	post, err := ctrl.service.GetPostByID(uint(id))

	if err != nil {
		return response.Error(c, http.StatusNotFound, "Post not found", err.Error())
	}

	res := dto.PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		Content:      post.Content,
		IsPublished:  post.IsPublished,
		Author:       post.Author.Username,
		PostCategory: post.Category.Title,
	}

	return response.Success(c, "Post retrieved successfully", res, 200)
}

func (ctrl *PostController) GetAllPosts(c echo.Context) error {
	page := 1
	limit := 10

	// Read query params
	if p := c.QueryParam("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if l := c.QueryParam("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if limit > 100 {
		limit = 100
	}

	posts, total, err := ctrl.service.GetAllPosts(page, limit)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to fetch posts", err.Error())
	}

	// Convert to DTO
	var postResponses []dto.PostResponse

	for _, p := range posts {
		postResponses = append(postResponses, dto.PostResponse{
			ID:           p.ID,
			Title:        p.Title,
			Content:      p.Content,
			IsPublished:  p.IsPublished,
			Author:       p.Author.Username,
			PostCategory: p.Category.Title,
		})
	}

	return response.Success(c, "Posts retrieved successfully", map[string]interface{}{
		"data":       postResponses,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	}, 200)
}

func (ctrl *PostController) UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid post ID", err.Error())
	}

	req := new(dto.UpdatePostRequest)

	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	post := models.Post{
		Title:          req.Title,
		Content:        req.Content,
		IsPublished:    req.IsPublished,
		PostCategoryID: req.PostCategoryID,
	}

	updatedPost, err := ctrl.service.UpdatePost(uint(id), &post)
	if err != nil {
		return response.Error(c, http.StatusNotFound, "Post not found", err.Error())
	}

	res := dto.PostResponse{
		ID:           updatedPost.ID,
		Title:        updatedPost.Title,
		Content:      updatedPost.Content,
		IsPublished:  updatedPost.IsPublished,
		Author:       updatedPost.Author.Username,
		PostCategory: updatedPost.Category.Title,
	}

	return response.Success(c, "Post updated successfully", res, 200)
}

func (ctrl *PostController) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid post ID", err.Error())
	}

	err = ctrl.service.DeletePost(uint(id))

	if err != nil {
		return response.Error(c, http.StatusNotFound, "Post not found", err.Error())
	}

	return response.Success(c, "Post deleted successfully", nil, 200)
}
