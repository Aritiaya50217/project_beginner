package marketdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
)

type FinnhubAdapter struct {
	apiKey string
}

type finnhubQuote struct {
	CurrentPrice float64 `json:"c"`
	HighPrice    float64 `json:"h"`
	LowPrice     float64 `json:"l"`
	OpenPrice    float64 `json:"o"`
	PrevClose    float64 `json:"pc"`
}

func NewFinnhubAdapter() port.MarketDataProvider {
	return &FinnhubAdapter{
		apiKey: os.Getenv("FINNHUB_API_KEY"),
	}
}

func (a *FinnhubAdapter) FetchQuote(symbol string) (*domain.StockQuote, error) {
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, a.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: status %d", resp.StatusCode)
	}

	var fq finnhubQuote
	if err := json.NewDecoder(resp.Body).Decode(&fq); err != nil {
		return nil, err
	}

	return &domain.StockQuote{
		Symbol:    symbol,
		Price:     fq.CurrentPrice,
		High:      fq.HighPrice,
		Low:       fq.LowPrice,
		Open:      fq.OpenPrice,
		PrevClose: fq.PrevClose,
	}, nil
}

func (a *FinnhubAdapter) FetchQuotes(symbols []string) ([]*domain.StockQuote, error) {
	quotes := []*domain.StockQuote{}

	for _, symbol := range symbols {
		quote, err := a.FetchQuote(symbol)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}
	return quotes, nil
}
