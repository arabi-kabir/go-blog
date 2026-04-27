package routes

import (
	"go-blog/controllers"
	"go-blog/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(
	e *echo.Echo,
	authController *controllers.AuthController,
	postController *controllers.PostController,
) {
	api := e.Group("/api")

	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)

	api.Use(middleware.JWTMiddleware())

	api.GET("/posts", postController.GetAllPosts)
	api.GET("/posts/:id", postController.GetPostByID)
	api.POST("/posts", postController.CreatePost)
	api.PUT("/posts/:id", postController.UpdatePost)
	api.DELETE("/posts/:id", postController.DeletePost)
}
