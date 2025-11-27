package domain

type ContractCreatedEvent struct {
	ContractID uint   `json:"contract_id"`
	UserID     uint   `json:"user_id"`
	Message    string `json:"message"`
}
