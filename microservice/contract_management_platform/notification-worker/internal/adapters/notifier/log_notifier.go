package notifier

import "fmt"

type LogNotifier struct{}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}

func (n *LogNotifier) Send(userID uint, message string) error {
	fmt.Printf("[NOTIFY] User %d -> %s\n", userID, message)
	return nil
}
