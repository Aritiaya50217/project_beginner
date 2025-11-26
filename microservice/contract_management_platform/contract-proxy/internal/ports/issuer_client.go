package ports

type IssuerClient interface {
	GetContracts(userID uint) ([]byte, error)
}
