/*
	Main gateway interface. Handles receiving transactions from the client and sending them to the processor.

	TODO: 
		- Integrate with a real processor
			- Establish connection
			- Send transaction
			- Receive response
			- Normalize Response
			- Handle Response
*/

package gateway

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"errors"
	"github.com/gbburleigh/quick-gateway/pkg/transaction"
	"github.com/gbburleigh/quick-gateway/pkg/types"
	"strconv"
	"strings"
)

func stringToTime(str string) (time.Time, error) {
	// Split the string by "/"
	parts := strings.Split(str, "/")
	if len(parts) != 2 {
			return time.Time{}, errors.New("invalid date format")
	}

	// Parse month and day
	month, err := strconv.Atoi(parts[0])
	if err != nil {
			return time.Time{}, fmt.Errorf("invalid month: %w", err)
	}
	year, err := strconv.Atoi(parts[1])
	if err != nil {
			return time.Time{}, fmt.Errorf("invalid day: %w", err)
	}

	firstDayOfMonth := time.Date(2000 + year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	return firstDayOfMonth, nil
}

type Gateway struct {
	ID string
}

func NewGateway() Gateway {
	return Gateway{
		ID: uuid.New().String(),
	}
}

func (g *Gateway) HandleTransaction(trx transaction.Transaction) error {
	if err := g.ValidateTransaction(trx); err != nil {
		fmt.Println("Transaction validation failed:", err)
		return err
	}

	g.SendToProcessor(trx)

	return nil
}

func (g *Gateway) SendToProcessor(trx transaction.Transaction) error {
	fmt.Println("Sending transaction to processor")
	return nil
}

func (g *Gateway) ReceiveFromProcessor(trx transaction.Transaction) error {
	fmt.Println("Received transaction from processor")
	return nil
}

func (g *Gateway) ValidateTransaction(trx transaction.Transaction) error {
	fmt.Println("Validating transaction", trx.Method)

	var Method = types.PaymentMethod(trx.Method)

	var Expiry, _ = stringToTime(Method.Expiry)

	if time.Now().After(Expiry) {
		return errors.New("Transaction expired")
	}

	if Method.CVV == "" {
		return errors.New("CVV is required")
	}

	if Method.HolderName == "" {
		return errors.New("Holder name is required")
	}

	if Method.Number == "" {
		return errors.New("Card number is required")
	}

	if trx.TransactionType == "card_present" && !trx.CardPresent {
		return errors.New("Card present is required")
	}

	if trx.TransactionType == "ach" && !trx.ACH {
		return errors.New("ACH is required")
	}

	if trx.TransactionType == "card_not_present" && trx.CardPresent {
		return errors.New("Card not present is required")
	}

	return nil
}

func (g *Gateway) NormalizeTransaction(trx transaction.Transaction) error {
	fmt.Println("Normalizing transaction")
	return nil
}

func (g *Gateway) LogTransaction(trx transaction.Transaction) error {
	fmt.Println("Logging transaction")
	return nil
}

func (g *Gateway) EstablishProcessorConnection() error {
	fmt.Println("Establishing processor connection")
	return nil
}
