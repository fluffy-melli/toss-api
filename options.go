package toss

type PaymentCancelOptionsRefundReceiveAccount struct {
	Bank          string
	AccountNumber string
	HolderName    string
}

type PaymentCancelOptions struct {
	CancelAmount         *int
	RefundReceiveAccount *PaymentCancelOptionsRefundReceiveAccount
	TaxFreeAmount        *int
	Currency             *string
}

type TransactionsOptions struct {
	startingAfter *string
	limit         *int
}
