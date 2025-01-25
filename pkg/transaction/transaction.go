package transaction
import (
	"github.com/gbburleigh/quick-gateway/pkg/types"
)

import (
	"github.com/google/uuid"
	"time"
)

// Transaction data
type Transaction struct {
	ID          string    `json:"id"`
	TransactionType        types.TransactionType    `json:"type"`
	Amount      float64   `json:"amount"`
	CardPresent bool      `json:"card_present"`
	ACH         bool      `json:"ach"`
	Method      types.PaymentMethod `json:"method"`
	Timestamp   time.Time `json:"timestamp"`
}



func NewTransaction(_type types.TransactionType, amount float64, cardPresent bool, ach bool, method types.PaymentMethod) Transaction {
	return Transaction{
		ID: uuid.New().String(),
		TransactionType: _type,
		Amount: amount,
		CardPresent: cardPresent,
		ACH: ach,
		Method: method,
		Timestamp: time.Now(),
	}
}


