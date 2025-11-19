package ports

import "contract-service/internal/domain"

type ContractRepository interface {
	Create(contract *domain.Contract) error
}
