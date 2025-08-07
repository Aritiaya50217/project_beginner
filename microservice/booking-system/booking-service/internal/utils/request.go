package utils

type BookingRequest struct {
	UserID    int    `json:"user_id"`
	ItemID    int    `json:"item_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type BookingResponse struct {
	ID        int          `json:"id"`
	UserID    UserResponse `json:"user_id"`
	ItemID    ItemResponse `json:"item_id"`
	StartTime string       `json:"start_time"`
	EndTime   string       `json:"end_time"`
	CreateAt  string       `json:"create_at"`
	UpdateAt  string       `json:"update_at"`
}

type ItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
