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

func (h *ItemHandler) GetAll(c *gin.Context) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	items, total, err := h.usecase.GetAllItem(c, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bookings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       items,
		"total":      total,
		"offset":     offset,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // คำนวณจำนวนหน้า
	})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
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
	err = h.usecase.DeleteItem(c, item.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted item is successfully"})
}
