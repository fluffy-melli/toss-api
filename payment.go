package toss

import (
	"time"

	"github.com/fluffy-melli/toss-api/bankcode"
	"github.com/fluffy-melli/toss-api/status"
	"github.com/fluffy-melli/toss-api/types"
)

type PaymentCancel struct {
	CancelAmount           int       `json:"cancelAmount"`
	CancelReason           string    `json:"cancelReason"`
	TaxFreeAmount          int       `json:"taxFreeAmount"`
	TaxExemptionAmount     int       `json:"taxExemptionAmount"`
	RefundableAmount       int       `json:"refundableAmount"`
	TransferDiscountAmount int       `json:"transferDiscountAmount"`
	EasyPayDiscountAmount  int       `json:"easyPayDiscountAmount"`
	CanceledAt             time.Time `json:"canceledAt"`
	TransactionKey         string    `json:"transactionKey"`
	ReceiptKey             string    `json:"receiptKey,omitempty"`
	CancelStatus           string    `json:"cancelStatus"`
	CancelRequestId        string    `json:"cancelRequestId"`
}

type PaymentCard struct {
	Amount                int                 `json:"amount"`
	IssuerCode            bankcode.Card       `json:"issuerCode"`
	AcquirerCode          bankcode.Card       `json:"acquirerCode,omitempty"`
	Number                string              `json:"number"`
	InstallmentPlanMonths int                 `json:"installmentPlanMonths"`
	ApproveNo             string              `json:"approveNo"`
	UseCardPoint          bool                `json:"useCardPoint"`
	CardType              types.Card          `json:"cardType"`
	OwnerType             types.Owner         `json:"ownerType"`
	AcquireStatus         status.Acquire      `json:"acquireStatus"`
	IsInterestFree        bool                `json:"isInterestFree"`
	InterestPayer         types.InterestPayer `json:"interestPayer,omitempty"`
}

type PaymentRefundReceiveAccount struct {
	BankCode      bankcode.Card `json:"bankCode"`
	AccountNumber string        `json:"accountNumber"`
	HolderName    string        `json:"holderName"`
}

type PaymentVirtualAccount struct {
	AccountType          types.Account               `json:"accountType"`
	AccountNumber        string                      `json:"accountNumber"`
	BankCode             bankcode.Card               `json:"bankCode"`
	CustomerName         string                      `json:"customerName"`
	DueDate              time.Time                   `json:"dueDate"`
	RefundStatus         status.Refund               `json:"refundStatus"`
	Expired              bool                        `json:"expired"`
	SettlementStatus     status.Settlement           `json:"settlementStatus"`
	RefundReceiveAccount PaymentRefundReceiveAccount `json:"refundReceiveAccount,omitempty"`
}

type PaymentMobilePhone struct {
	CustomerMobilePhone string            `json:"customerMobilePhone"`
	SettlementStatus    status.Settlement `json:"settlementStatus"`
	ReceiptUrl          string            `json:"receiptUrl"`
}

type PaymentGiftCertificate struct {
	ApproveNo        string            `json:"approveNo"`
	SettlementStatus status.Settlement `json:"settlementStatus"`
}

type PaymentTransfer struct {
	BankCode         bankcode.Card     `json:"bankCode"`
	SettlementStatus status.Settlement `json:"settlementStatus"`
}

type PaymentEasyPay struct {
	Provider       bankcode.EasyPay `json:"provider"`
	Amount         int              `json:"amount"`
	DiscountAmount int              `json:"discountAmount"`
}

type PaymentURL struct {
	URL string `json:"url"`
}

type PaymentFailure struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaymentCashReceipt struct {
	Type          string `json:"type"`
	ReceiptKey    string `json:"receiptKey"`
	IssueNumber   string `json:"issueNumber"`
	ReceiptUrl    string `json:"receiptUrl"`
	Amount        int    `json:"amount"`
	TaxFreeAmount int    `json:"taxFreeAmount"`
}

type PaymentCashReceipts struct {
	ReceiptKey             string          `json:"receiptKey"`
	OrderId                string          `json:"orderId"`
	OrderName              string          `json:"orderName"`
	Type                   string          `json:"type"`
	IssueNumber            string          `json:"issueNumber"`
	ReceiptUrl             string          `json:"receiptUrl"`
	BusinessNumber         string          `json:"businessNumber"`
	TransactionType        string          `json:"transactionType"`
	Amount                 int             `json:"amount"`
	TaxFreeAmount          int             `json:"taxFreeAmount"`
	IssueStatus            string          `json:"issueStatus"`
	Failure                *PaymentFailure `json:"failure,omitempty"`
	CustomerIdentityNumber string          `json:"customerIdentityNumber"`
	RequestedAt            time.Time       `json:"requestedAt"`
}

type PaymentDiscount struct {
	Amount int `json:"amount"`
}

type Payment struct {
	Version             string                  `json:"version"`
	PaymentKey          string                  `json:"paymentKey"`
	Type                types.Payment           `json:"type"`
	OrderId             string                  `json:"orderId"`
	OrderName           string                  `json:"orderName"`
	MID                 string                  `json:"mId"`
	Currency            string                  `json:"currency"`
	Method              types.Method            `json:"method,omitempty"`
	TotalAmount         int                     `json:"totalAmount"`
	BalanceAmount       int                     `json:"balanceAmount"`
	Status              status.Payment          `json:"status"`
	RequestedAt         time.Time               `json:"requestedAt"`
	ApprovedAt          *time.Time              `json:"approvedAt,omitempty"`
	UseEscrow           bool                    `json:"useEscrow"`
	LastTransactionKey  string                  `json:"lastTransactionKey,omitempty"`
	SuppliedAmount      int                     `json:"suppliedAmount"`
	Vat                 int                     `json:"vat"`
	CultureExpense      bool                    `json:"cultureExpense"`
	TaxFreeAmount       int                     `json:"taxFreeAmount"`
	TaxExemptionAmount  int                     `json:"taxExemptionAmount"`
	Cancels             []PaymentCancel         `json:"cancels,omitempty"`
	IsPartialCancelable bool                    `json:"isPartialCancelable"`
	Card                *PaymentCard            `json:"card,omitempty"`
	VirtualAccount      *PaymentVirtualAccount  `json:"virtualAccount,omitempty"`
	Secret              string                  `json:"secret,omitempty"`
	MobilePhone         *PaymentMobilePhone     `json:"mobilePhone,omitempty"`
	GiftCertificate     *PaymentGiftCertificate `json:"giftCertificate,omitempty"`
	Transfer            *PaymentTransfer        `json:"transfer,omitempty"`
	Metadata            map[string]any          `json:"metadata,omitempty"`
	Receipt             *PaymentURL             `json:"receipt,omitempty"`
	Checkout            *PaymentURL             `json:"checkout,omitempty"`
	EasyPay             *PaymentEasyPay         `json:"easyPay,omitempty"`
	Country             string                  `json:"country"`
	Failure             *PaymentFailure         `json:"failure,omitempty"`
	CashReceipt         *PaymentCashReceipt     `json:"cashReceipt,omitempty"`
	CashReceipts        []PaymentCashReceipts   `json:"cashReceipts,omitempty"`
	Discount            *PaymentDiscount        `json:"discount,omitempty"`
}
