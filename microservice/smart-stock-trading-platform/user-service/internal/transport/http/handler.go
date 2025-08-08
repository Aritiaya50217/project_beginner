package http

import (
	"net/http"
	"smart-stock-trading-platform-user-service/internal/usecase"
	"smart-stock-trading-platform-user-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Routers(r *gin.RouterGroup) {
	r.POST("/register", h.register)
}

func (h *UserHandler) register(c *gin.Context) {
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
