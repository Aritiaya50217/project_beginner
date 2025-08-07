package handler

import (
	"booking-system-booking-service/internal/app"
	"booking-system-booking-service/internal/domain"
	"booking-system-booking-service/internal/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	usecase app.BookingUsecase
}

func NewBookingHandler(usecase app.BookingUsecase) *BookingHandler {
	return &BookingHandler{usecase: usecase}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	// ดึง Authorization header มา
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	// Get user ID from JWT token
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

	var req utils.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	booking := domain.Booking{
		UserID:    uint(userIDInt),
		ItemID:    uint(req.ItemID),
		StartTime: startTime,
		EndTime:   endTime,
	}

	if err := h.usecase.Create(c, &booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "booking created successfully"})
}

func (h *BookingHandler) GetByID(c *gin.Context) {
	// ดึง Authorization header มา
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

	// ดึง booking ID
	bookingIDParam := c.Param("id")
	bookingID, err := strconv.Atoi(bookingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	booking, err := h.usecase.GetBookingByID(c.Request.Context(), bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if booking.UserID != uint(userIDInt) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	userIDStr := strconv.FormatUint(uint64(booking.UserID), 10)

	user, err := utils.GetUserInfo(os.Getenv("USER_SERVICE_URL"), userIDStr, token)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch user info"})
		return
	}

	var userInfo utils.UserResponse
	userInfo.ID = user.ID
	userInfo.Name = user.Name
	userInfo.Email = user.Email

	var itemInfo utils.ItemResponse
	itemInfo.ID = booking.Item.ID
	itemInfo.Name = booking.Item.Name

	result := utils.BookingResponse{
		ID:        booking.ID,
		UserID:    userInfo,
		ItemID:    itemInfo,
		StartTime: utils.FormatDate(booking.StartTime),
		EndTime:   utils.FormatDate(booking.EndTime),
		CreateAt:  utils.FormatDate(booking.CreatedAt),
		UpdateAt:  utils.FormatDate(booking.UpdatedAt),
	}

	c.JSON(http.StatusOK, result)
}

func (h *BookingHandler) GetAll(c *gin.Context) {
	// ดึง Authorization header มา
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

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

	bookings, total, err := h.usecase.GetAllBooking(c, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get bookings"})
		return
	}


	for _, booking := range bookings {
		if booking.UserID != uint(userIDInt) {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}

		userIDStr := strconv.FormatUint(uint64(booking.UserID), 10)

		user, err := utils.GetUserInfo(os.Getenv("USER_SERVICE_URL"), userIDStr, token)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch user info"})
			return
		}

		if uint(user.ID) != booking.UserID {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       bookings,
		"total":      total,
		"offset":     offset,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // คำนวณจำนวนหน้า
	})

}
