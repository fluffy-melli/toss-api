package toss

import (
	"time"

	"github.com/fluffy-melli/toss-api/status"
	"github.com/fluffy-melli/toss-api/types"
)

type Transaction struct {
	MID            string         `json:"mId"`
	TransactionKey string         `json:"transactionKey"`
	PaymentKey     string         `json:"paymentKey"`
	OrderId        string         `json:"orderId"`
	Method         types.Method   `json:"method"`
	CustomerKey    string         `json:"customerKey"`
	UseEscrow      bool           `json:"useEscrow"`
	ReceiptUrl     string         `json:"receiptUrl"`
	Status         status.Payment `json:"status"`
	TransactionAt  time.Time      `json:"transactionAt"`
	Currency       string         `json:"currency"`
	Amount         int            `json:"amount"`
}
