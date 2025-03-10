package handler

import (
	"ECCO2K/internal/models"
	"ECCO2K/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{"error": "неккоректные данные"})
		return
	}

	if err := h.service.CreateUser(context.Background(), &user); err != nil {
		c.JSON(201, gin.H{"error": "ошибка при создании пользователя"})
		return
	}

	c.JSON(201, gin.H{"user": user})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(404, gin.H{"error": "неккоректные данные"})
		return
	}
	user, err := h.service.GetUserByID(context.Background(), uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "неккоректные данные"})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) Register(c *gin.Context) {
	var request struct {
		Name     string `json:"Name"`
		Password string `json:"Password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(404, gin.H{"error": "неккоректные данные"})
		return
	}
	user, err := h.service.Register(context.Background(), request.Name, request.Password)
	if err != nil {
		c.JSON(404, gin.H{"error": "ошибка при регистрации"})
		return
	}
}
