package handler

import (
	"booking-system-booking-service/internal/app"
	"booking-system-booking-service/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	usecase app.ItemUsecase
}

func NewItemHandler(u app.ItemUsecase) *ItemHandler {
	return &ItemHandler{usecase: u}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Create(c, &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) GetByID(c *gin.Context) {
	itemIDParam := c.Param("id")
	itemID, err := strconv.Atoi(itemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	item, err := h.usecase.GetByID(c, itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": item})

}
