package main

import (
	"go-blog/config"
	"go-blog/controllers"
	"go-blog/middleware"
	"go-blog/repositories"
	"go-blog/routes"
	"go-blog/services"

	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Connect to the database
	db := config.ConnectDB()

	// setup middleware
	middleware.SetupMiddleware(e)

	// Wire DI chain: DB → Repository → Service → Controller
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	authService := services.NewAuthService(userRepo)
	postService := services.NewPostService(postRepo)

	authController := controllers.NewAuthController(authService)
	postController := controllers.NewPostController(postService)

	// Register routes
	routes.InitRoutes(e, authController, postController)

	// Start Server
	port := ":8080"
	log.Println("🚀 Server running on http://localhost" + port)

	if err := e.Start(port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
