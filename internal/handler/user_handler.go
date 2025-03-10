package handler

import (
	"ECCO2K/internal/models"
	"ECCO2K/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewUserHandler(userService *service.UserService, authService *service.AuthService) *UserHandler {
	return &UserHandler{userService: userService, authService: authService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "некорректные данные"})
		return
	}

	if err := h.userService.CreateUser(context.Background(), &user); err != nil {
		c.JSON(500, gin.H{"error": "ошибка при создании пользователя"})
		return
	}

	c.JSON(201, gin.H{"user": user})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "некорректные данные"})
		return
	}

	user, err := h.userService.GetUserByID(context.Background(), uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "пользователь не найден"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) Register(c *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "некорректные данные"})
		return
	}

	user, err := h.authService.Register(context.Background(), request.Name, request.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "некорректные данные"})
		return
	}

	token, err := h.authService.Login(context.Background(), request.Name, request.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
