package ports

type Notifier interface {
	Send(userID uint, message string) error
}
