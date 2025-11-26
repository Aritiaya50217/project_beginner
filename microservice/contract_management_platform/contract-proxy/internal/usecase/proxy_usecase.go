package usecase

import "contract-proxy/internal/ports"

type ProxyUsecase struct {
	issuerClient ports.IssuerClient
}

func NewProxyUsecase(client ports.IssuerClient) *ProxyUsecase {
	return &ProxyUsecase{issuerClient: client}
}

func (uc *ProxyUsecase) GetContracts(userID uint) ([]byte, error) {
	return uc.issuerClient.GetContracts(userID)
}
