package utils

// ข้อมูล quote ที่มาจาก API ภายนอก (ไม่จำเป็นต้องบันทึก DB)
type StockQuote struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	PrevClose float64 `json:"prevClose"`
}

type Company struct {
	Country              string  `json:"country"`
	Currency             string  `json:"currency"`
	Exchange             string  `json:"exchange"`
	Ipo                  string  `json:"ipo"`
	MarketCapitalization float32 `json:"marketCapitalization"`
	Name                 string  `json:"name"`
	Phone                string  `json:"phone"`
	ShareOutstanding     float32 `json:"shareOutstanding"`
	Ticker               string  `json:"ticker"`
	Weburl               string  `json:"weburl"`
	Logo                 string  `json:"logo"`
	FinnhubIndustry      string  `json:"finnhubIndustry"`
}
