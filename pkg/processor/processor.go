/*
	Main processor interface. Handles receiving transactions from the gateway and processing them into deposits.
	

	I'm not sure if we need this yet but it's stubbed out here for future use.
*/
package processor

import (
	"fmt"
	"github.com/gbburleigh/quick-gateway/pkg/transaction"
	"github.com/google/uuid"
)

type Processor struct {
	ID string
}

func NewProcessor() Processor {
	return Processor{
		ID: uuid.New().String(),
	}
}

func (p *Processor) ProcessTransaction(transaction transaction.Transaction) error {
	fmt.Println("Processing transaction")
	return nil
}
