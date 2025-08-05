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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userIDFormToken, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if int(userIDFormToken.(float64)) != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this user"})
		return
	}

	var req utils.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get existing user from DB
	user, err := h.usecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.Name = req.Name
	user.Email = req.Email

	if err = h.usecase.UpdateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}
