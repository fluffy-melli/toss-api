package toss

import "time"

type Transaction struct {
	MID            string    `json:"mId"`
	TransactionKey string    `json:"transactionKey"`
	PaymentKey     string    `json:"paymentKey"`
	OrderId        string    `json:"orderId"`
	Method         Method    `json:"method"`
	CustomerKey    string    `json:"customerKey"`
	UseEscrow      bool      `json:"useEscrow"`
	ReceiptUrl     string    `json:"receiptUrl"`
	Status         Status    `json:"status"`
	TransactionAt  time.Time `json:"transactionAt"`
	Currency       string    `json:"currency"`
	Amount         int       `json:"amount"`
}
