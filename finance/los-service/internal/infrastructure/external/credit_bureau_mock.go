package external

import (
	"log"
)

func CheckCreditScore(customerID int64) int {
	log.Printf("Mock checking credit score for customer %d\n", customerID)
	return 700
}
