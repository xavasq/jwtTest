package router

import (
	"ECCO2K/internal/config"
	"ECCO2K/internal/handler"
	"ECCO2K/internal/middlewares"
	"ECCO2K/internal/repository"
	"ECCO2K/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RouterSetup(r *gin.Engine, db *pgxpool.Pool) {
	cfg := config.LoadEnv()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, cfg.JWT_SECRET)
	userHandler := handler.NewUserHandler(userService, authService)

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/user/:id", userHandler.GetUserByID)
	}

	private := api.Group("/private")
	private.Use(middlewares.AuthMiddleware())
	{
		private.GET("/profile", userHandler.GetUserByID)
	}
}
