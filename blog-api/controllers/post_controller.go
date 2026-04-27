package controllers

import (
	"go-blog/models"
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	createdPost, err := ctrl.service.CreatePost(post)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create post",
		})
	}
	return c.JSON(http.StatusOK, createdPost)
}

func (ctrl *PostController) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid post ID",
		})
	}

	post, err := ctrl.service.GetPostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Post not found",
		})
	}
	return c.JSON(http.StatusOK, post)
}

func (ctrl *PostController) GetAllPosts(c echo.Context) error {
	posts, err := ctrl.service.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch posts",
		})
	}
	return c.JSON(http.StatusOK, posts)
}

func (ctrl *PostController) UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid post ID",
		})
	}

	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	updatedPost, err := ctrl.service.UpdatePost(uint(id), post)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Post not found",
		})
	}
	return c.JSON(http.StatusOK, updatedPost)
}

func (ctrl *PostController) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid post ID",
		})
	}

	err = ctrl.service.DeletePost(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Post not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Post deleted successfully",
	})
}
