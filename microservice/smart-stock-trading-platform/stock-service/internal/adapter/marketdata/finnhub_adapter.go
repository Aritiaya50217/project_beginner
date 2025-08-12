package marketdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
	"time"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
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
	limiter := time.Tick(200 * time.Millisecond) // 5 requests per second

	for _, symbol := range symbols {
		<-limiter // รอ tick ก่อนยิง request

		quote, err := a.FetchQuote(symbol)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}
	return quotes, nil
}

// ดึงรายชื่อหุ้นจาก exchange เช่น "US" หรือ "NASDAQ"
func (a *FinnhubAdapter) FetchSymbolList(exchange string) ([]string, error) {
	url := fmt.Sprintf("https://finnhub.io/api/v1/stock/symbol?exchange=%s&token=%s", exchange, a.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch symbols: %s", resp.Status)
	}

	var data []struct {
		Symbol string `json:"symbol"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	symbols := []string{}
	for _, val := range data {
		symbols = append(symbols, val.Symbol)
	}
	return symbols, nil
}

func (u *FinnhubAdapter) FetchCompayny(symbol string) (*domain.Company, error) {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", u.apiKey)
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	company, _, err := finnhubClient.CompanyProfile2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, err
	}

	data := domain.Company{
		Country:              *company.Country,
		Currency:             *company.Currency,
		Exchange:             *company.Exchange,
		Ipo:                  *company.Ipo,
		MarketCapitalization: *company.MarketCapitalization,
		Name:                 *company.Name,
		Phone:                *company.Phone,
		ShareOutstanding:     *company.ShareOutstanding,
		Ticker:               *company.Ticker,
		Weburl:               *company.Weburl,
		Logo:                 *company.Logo,
		FinnhubIndustry:      *company.FinnhubIndustry,
	}

	return &data, nil
}
