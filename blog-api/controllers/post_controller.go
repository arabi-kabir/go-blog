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

// PostController methods
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
	posts, err := ctrl.service.GetAllPosts()

	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to fetch posts", err.Error())
	}

	return response.Success(c, "Posts retrieved successfully", posts, 200)
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
