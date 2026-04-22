package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gummymule/task-manager/config"
	_ "github.com/gummymule/task-manager/docs"
	"github.com/gummymule/task-manager/internal/handler"
	"github.com/gummymule/task-manager/internal/repository"
	"github.com/gummymule/task-manager/internal/usecase"
	"github.com/gummymule/task-manager/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task Manager API
// @version         1.0
// @description     A RESTful API for task management built with Go and Clean Architecture.

// @contact.name    gummymule
// @contact.url     https://github.com/gummymule

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load config
	cfg := config.Load()

	// Init database
	db := cfg.InitDB()
	defer db.Close()

	// Init layers (repository → usecase → handler)
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	boardRepo := repository.NewBoardRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	boardHandler := handler.NewBoardHandler(boardUsecase)

	// Init router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	api := r.Group("/api/v1")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	// Protected routes
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		auth.POST("/logout", userHandler.Logout)

		// Board routes
		auth.GET("/boards", boardHandler.GetAll)
		auth.GET("/boards/:board_id", boardHandler.GetByID)
		auth.POST("/boards", boardHandler.Create)
		auth.PUT("/boards/:board_id", boardHandler.Update)
		auth.DELETE("/boards/:board_id", boardHandler.Delete)

		// Task routes
		boards := auth.Group("/boards/:board_id")
		{
			boards.GET("/tasks", taskHandler.GetAll)
			boards.POST("/tasks", taskHandler.Create)
		}

		tasks := auth.Group("/tasks")
		{
			tasks.GET("/:id", taskHandler.GetByID)
			tasks.PUT("/:id", taskHandler.Update)
			tasks.DELETE("/:id", taskHandler.Delete)
		}
	}

	// Swagger routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
