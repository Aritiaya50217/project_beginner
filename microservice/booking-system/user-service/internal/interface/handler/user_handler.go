package handler

import (
	"booking-system-user-service/internal/app"
	"booking-system-user-service/internal/domain"
	"booking-system-user-service/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase app.UserUsecase
}

func NewUserHandler(usecase app.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req utils.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	if err := h.usecase.Register(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register is successfuly"})

}

func (h *UserHandler) Login(c *gin.Context) {
	var req utils.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	user, err := h.usecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})

}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.usecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	var result []utils.UserResponse
	for _, user := range users {
		result = append(result, utils.UserResponse{
			ID:    int(user.ID),
			Email: user.Email,
			Name:  user.Name,
		})
	}
	c.JSON(http.StatusOK, result)
}
