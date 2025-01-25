# Quick Gateway

Quick Gateway is a simple gateway for processing payments. It is designed to be easy to use and integrate with existing systems.

## Getting Started

### Installation

```bash
go get github.com/quick-gateway/quick-gateway
```


### API

`quick-gateway` relies on a client-gateway-processor architecture. Transactions are created at the client level, to emulate a POS system such as a terminal. The client is responsible for creating the transaction and sending it to the gateway. The gateway is responsible for processing the transaction, normalizing and logging it as necessary, and sending it to the processor. The processor is responsible for processing the transaction and returning the result to the gateway. Most processors use proprietary APIs and require subscriptions to integrate with them, so `quick-gateway` is designed with a simple mock processor that simulates these services for us.

### Usage

#### Create a Client

```go
client := client.NewClient("terminal_id", "secret_key", "environment", "gateway_id")

// TODO: Hook into POS to read transaction data and trigger client to create transaction
const transaction_data = <Read from POS, API, etc>

const Transaction = client.CreateTransaction(
    transaction_data.Type,
    transaction_data.Amount,
    transaction_data.CardPresent,
    transaction_data.ACH,
    transaction_data.Method,
)

client.SendTransaction(Transaction)
```

#### Receive a Transaction

```go

const Transaction = <Transaction from Client>

const Gateway = gateway.NewGateway(client)

Gateway.HandleTransaction(Transaction)
```

