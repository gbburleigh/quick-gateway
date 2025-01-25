package main

import (
	"fmt"
	"github.com/gbburleigh/quick-gateway/pkg/gateway"
	"github.com/gbburleigh/quick-gateway/pkg/processor"
	"github.com/gbburleigh/quick-gateway/pkg/client"
	"github.com/gbburleigh/quick-gateway/pkg/transaction"
	"github.com/gbburleigh/quick-gateway/pkg/types"
)

func main() {
	fmt.Println()
	fmt.Println("quick-gateway")

	Gateway := gateway.NewGateway()
	Processor := processor.NewProcessor()
	Client := client.NewClient("terminal_id", "secret_key", "environment", Gateway.ID)

	fmt.Println()
	fmt.Println("Instantiated Gateway, Processor, and Client")
	fmt.Println("Gateway ID:", Gateway.ID)
	fmt.Println("Processor ID:", Processor.ID)
	fmt.Println("Client ID:", Client.ID)
	fmt.Println()

	var transaction = transaction.NewTransaction(
		"card_present",
		100.00,
		true,
		false,
		types.PaymentMethod{
			MethodType: "card",
			HolderName: "John Doe",
			Number: "4111111111111111",
			Expiry: "12/25",
			CVV: "123",
		},
	)

	var Transaction, err = Client.CreateTransaction(transaction.TransactionType, transaction.Amount, transaction.CardPresent, transaction.ACH, transaction.Method)

	if err != nil {
		fmt.Println("Transaction creation failed:", err)
		return
	}

	fmt.Println("Transaction created:", Transaction)

	var GatewayErrors = Gateway.HandleTransaction(Transaction)

	if GatewayErrors != nil {
		fmt.Println("Got Error from Gateway:", GatewayErrors)
		return
	}

	fmt.Println("Transaction processed successfully")
	fmt.Println()

	Transaction.Method.Expiry = "12/24"

	var ExpiredTransaction = Gateway.HandleTransaction(Transaction)

	fmt.Println("Expired Transaction:", ExpiredTransaction)
}
