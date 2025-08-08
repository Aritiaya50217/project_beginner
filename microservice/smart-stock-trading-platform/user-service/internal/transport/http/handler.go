package http

import (
	"log"
	"net/http"
	"smart-stock-trading-platform-user-service/internal/port"
	"smart-stock-trading-platform-user-service/internal/usecase"
	"smart-stock-trading-platform-user-service/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase     *usecase.UserUsecase
	authService port.AuthService
}

// NewUserHandler รับ usecase และ auth service (เช่น jwt)
func NewUserHandler(usecase *usecase.UserUsecase, auth port.AuthService) *UserHandler {
	return &UserHandler{
		usecase:     usecase,
		authService: auth,
	}
}
func (h *UserHandler) Register(c *gin.Context) {
	var req utils.ReqRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.Register(c, req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req utils.ReqLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// ดึง userID จาก context (เซ็ตโดย middleware)
	userIDFromToken, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// เช็คว่า userID ใน token ต้องตรงกับ param id
	if userIDFromToken.(uint) != uint(id) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: you can only access your own user data"})
		return
	}

	user, err := h.usecase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Printf("Returning user: %+v", user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetUpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// ดึง userID จาก context (เซ็ตโดย middleware)
	userIDFromToken, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// เช็คว่า userID ใน token ต้องตรงกับ param id
	if userIDFromToken.(uint) != uint(id) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: you can only access your own user data"})
		return
	}

	var req utils.ReqRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get existing user from DB
	user, err := h.usecase.GetUserByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName

	if err = h.usecase.UpdateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := h.usecase.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
