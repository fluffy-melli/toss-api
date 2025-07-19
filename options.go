package toss

import (
	"time"

	"github.com/fluffy-melli/toss-api/bankcode"
	"github.com/fluffy-melli/toss-api/types"
)

type PaymentCancelOptionsRefundReceiveAccount struct {
	Bank          *bankcode.Card
	AccountNumber *string
	HolderName    *string
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

type VirtualAccountsOptionsCashReceipt struct {
	Type               *types.CashReceipt
	RegistrationNumber *string
}

type VirtualAccountsEscrowProducts struct {
	ID        string
	Name      string
	Code      string
	UnitPrice int
	Quantity  int
}

type VirtualAccountsOptions struct {
	ValidHours          *int
	DueDate             *time.Time
	CustomerEmail       *string
	CustomerMobilePhone *string
	TaxFreeAmount       *int
	UseEscrow           *bool
	CashReceipt         *VirtualAccountsOptionsCashReceipt
}
