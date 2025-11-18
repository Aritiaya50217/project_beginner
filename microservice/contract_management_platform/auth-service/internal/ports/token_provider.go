package ports

type TokenProvider interface {
	ValidateToken(token string) (map[string]interface{}, error)
}
