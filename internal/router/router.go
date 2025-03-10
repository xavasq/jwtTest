package router

import (
	"ECCO2K/internal/handler"
	"ECCO2K/internal/repository"
	"ECCO2K/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RouterSetup(r *gin.Engine, db *pgxpool.Pool) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/user/:id", userHandler.GetUserByID)
	}
}
