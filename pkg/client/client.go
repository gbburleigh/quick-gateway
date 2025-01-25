/*
	Basic client interface for sending transactions to gateway. Gateway expects the following fields to create a transaction:

	@param Type 
*/

package client

import (
	"github.com/google/uuid"
	"github.com/gbburleigh/quick-gateway/pkg/transaction"
	"github.com/gbburleigh/quick-gateway/pkg/types"
	"fmt"
	"time"
	"errors"
)

type Client struct {
	ID          string    `json:"id"`
	TerminalId  string    `json:"terminal_id"`
	SecretKey   string    `json:"secret_key"`
	Environment string    `json:"environment"`
	GatewayId   string    `json:"gateway_id"`
}

func NewClient(terminalId, secretKey, environment string, gatewayId string) Client {
	return Client{
		ID:          uuid.New().String(),
		TerminalId:  terminalId,
		SecretKey:   secretKey,
		Environment: environment,
		GatewayId:   gatewayId,
	}
}

func (c *Client) SendTransaction(transaction transaction.Transaction) error {
	fmt.Println("Sending transaction to gateway")
	return nil
}

func (c *Client) CreateTransaction(_type types.TransactionType, amount float64, cardPresent bool, ach bool, method types.PaymentMethod) (transaction.Transaction, error) {
	var Transaction = transaction.Transaction{
		ID:          uuid.New().String(),
		TransactionType:        _type,
		Amount:      amount,
		CardPresent: cardPresent,
		ACH:         ach,
		Method:      method,
		Timestamp:   time.Now(),
	}

	if (Transaction == transaction.Transaction{}) {
		return transaction.Transaction{}, errors.New("Transaction Not Created")
	}

	return Transaction, nil
}

func (c *Client) ConnectToGateway(gatewayId string) error {
	fmt.Println("Connecting to gateway")
	return nil
}
