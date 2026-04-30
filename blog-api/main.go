package main

import (
	"go-blog/config"
	"go-blog/controllers"
	"go-blog/middleware"
	"go-blog/pkg/cache"
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

	// connct redis
	config.ConnectRedis()

	// setup middleware
	middleware.SetupMiddleware(e)

	cacheLayer := cache.NewCache(config.RDB, config.Ctx)

	// Wire DI chain: DB → Repository → Service → Controller
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	authService := services.NewAuthService(userRepo)
	postService := services.NewPostService(postRepo, userRepo, cacheLayer)

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
