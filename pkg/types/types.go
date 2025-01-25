package types

type PaymentMethod struct {
	MethodType MethodType `json:"method_type"` 
	HolderName string      `json:"holder_name"` // Name on account
	Number     string      `json:"number"`
	Expiry     string      `json:"expiry"`
	CVV        string      `json:"cvv"` // CP/CNP specific
	AccountNumber string `json:"account_number"` // ACH specific
	RoutingNumber string `json:"routing_number"`
}

type TransactionType string

const (
	Charge TransactionType    = "charge"
	Refund TransactionType    = "refund"
	Chargeback TransactionType = "chargeback"
	Void  TransactionType = "void"
)

type MethodType string

const (
	CardPresent MethodType = "card_present"
	ACH         MethodType = "ach"
	CardNotPresent MethodType = "card_not_present"
)