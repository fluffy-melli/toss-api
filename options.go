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

type TransactionOptions struct {
	StartingAfter *string
	Limit         *int
}

type SettlementOptions struct {
	Page *int
	Size *int
}

type BillingOptions struct {
	CustomerEmail      *string
	CustomerName       *string
	TaxFreeAmount      *int
	TaxExemptionAmount *int
}
