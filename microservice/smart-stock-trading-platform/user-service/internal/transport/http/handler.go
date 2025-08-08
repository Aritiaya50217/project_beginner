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

	token, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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
