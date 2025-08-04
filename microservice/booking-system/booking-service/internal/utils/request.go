package utils

type BookingRequest struct {
	UserID    int    `json:"user_id"`
	ItemID    int    `json:"item_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
