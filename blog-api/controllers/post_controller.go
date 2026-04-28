package controllers

import (
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
	post := new(models.Post)

	if err := c.Bind(post); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	createdPost, err := ctrl.service.CreatePost(post)

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to create post", err.Error())
	}

	return response.Success(c, "Post created successfully", createdPost, 201)
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

	return response.Success(c, "Post retrieved successfully", post, 200)
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

	return response.Success(c, "Posts retrieved successfully", map[string]interface{}{
		"data":       posts,
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

	post := new(models.Post)

	if err := c.Bind(post); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	updatedPost, err := ctrl.service.UpdatePost(uint(id), post)
	if err != nil {
		return response.Error(c, http.StatusNotFound, "Post not found", err.Error())
	}

	return response.Success(c, "Post updated successfully", updatedPost, 200)
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
