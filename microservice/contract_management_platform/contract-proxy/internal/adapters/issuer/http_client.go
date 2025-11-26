package issuer

import (
	"contract-proxy/internal/ports"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewHTTPClient(baseURL string) ports.IssuerClient {
	return &HTTPClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (h *HTTPClient) GetContracts(userID uint) ([]byte, error) {
	url := fmt.Sprintf("%s/contracts/%d", h.baseURL, userID)
	resp, err := h.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
