package domain

type Item struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
