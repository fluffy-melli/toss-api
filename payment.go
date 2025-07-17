package toss

import (
	"time"
)

type PaymentType string

const (
	PaymentTypeNormal   PaymentType = "NORMAL"
	PaymentTypeBilling  PaymentType = "BILLING"
	PaymentTypeBrandpay PaymentType = "BRANDPAY"
)

type Method string

const (
	MethodCard           Method = "카드"
	MethodVirtualAccount Method = "가상계좌"
	MethodEasyPay        Method = "간편결제"
	MethodMobilePhone    Method = "핸드폰"
	MethodTransfer       Method = "계좌이체"
	MethodCulture        Method = "문화상품권"
	MethodBookCulture    Method = "도서문화상품권"
	MethodGameCulture    Method = "게임문화상품권"
)

type Status string

const (
	StatusReady             Status = "READY"
	StatusInProgress        Status = "IN_PROGRESS"
	StatusWaitingForDeposit Status = "WAITING_FOR_DEPOSIT"
	StatusDone              Status = "DONE"
	StatusCanceled          Status = "CANCELED"
	StatusPartialCanceled   Status = "PARTIAL_CANCELED"
	StatusAborted           Status = "ABORTED"
	StatusExpired           Status = "EXPIRED"
)

type PaymentCancels struct {
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

type PaymentCardType string

const (
	PaymentCardTypeCredit  PaymentCardType = "신용"
	PaymentCardTypeDebit   PaymentCardType = "체크"
	PaymentCardTypeGift    PaymentCardType = "기프트"
	PaymentCardTypeUnknown PaymentCardType = "미확인"
)

type PaymentOwnerType string

const (
	PaymentOwnerTypePersonal  PaymentOwnerType = "개인"
	PaymentOwnerTypeCorporate PaymentOwnerType = "법인"
	PaymentOwnerTypeUnknown   PaymentOwnerType = "미확인"
)

type PaymentAcquireStatus string

const (
	PaymentAcquireStatusReady           PaymentAcquireStatus = "READY"
	PaymentAcquireStatusRequested       PaymentAcquireStatus = "REQUESTED"
	PaymentAcquireStatusCompleted       PaymentAcquireStatus = "COMPLETED"
	PaymentAcquireStatusCancelRequested PaymentAcquireStatus = "CANCEL_REQUESTED"
	PaymentAcquireStatusCanceled        PaymentAcquireStatus = "CANCELED"
)

type PaymentInterestPayer string

const (
	PaymentInterestPayerBuyer       PaymentInterestPayer = "BUYER"
	PaymentInterestPayerCardCompany PaymentInterestPayer = "CARD_COMPANY"
	PaymentInterestPayerMerchant    PaymentInterestPayer = "MERCHANT"
)

type PaymentCard struct {
	Amount                int                  `json:"amount"`
	IssuerCode            string               `json:"issuerCode"`
	AcquirerCode          string               `json:"acquirerCode,omitempty"`
	Number                string               `json:"number"`
	InstallmentPlanMonths int                  `json:"installmentPlanMonths"`
	ApproveNo             string               `json:"approveNo"`
	UseCardPoint          bool                 `json:"useCardPoint"`
	CardType              PaymentCardType      `json:"cardType"`
	OwnerType             PaymentOwnerType     `json:"ownerType"`
	AcquireStatus         PaymentAcquireStatus `json:"acquireStatus"`
	IsInterestFree        bool                 `json:"isInterestFree"`
	InterestPayer         PaymentInterestPayer `json:"interestPayer,omitempty"`
}

type PaymentAccountType string

const (
	PaymentAccountTypeGeneral PaymentAccountType = "일반"
	PaymentAccountTypeFixed   PaymentAccountType = "고정"
)

type PaymentRefundStatus string

const (
	PaymentRefundStatusNone          PaymentRefundStatus = "NONE"
	PaymentRefundStatusPending       PaymentRefundStatus = "PENDING"
	PaymentRefundStatusFailed        PaymentRefundStatus = "FAILED"
	PaymentRefundStatusPartialFailed PaymentRefundStatus = "PARTIAL_FAILED"
	PaymentRefundStatusCompleted     PaymentRefundStatus = "COMPLETED"
)

type PaymentSettlementStatus string

const (
	PaymentSettlementStatusIncompleted PaymentSettlementStatus = "INCOMPLETED"
	PaymentSettlementStatusCompleted   PaymentSettlementStatus = "COMPLETED"
)

type PaymentRefundReceiveAccount struct {
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	HolderName    string `json:"holderName"`
}

type PaymentVirtualAccount struct {
	AccountType          PaymentAccountType          `json:"accountType"`
	AccountNumber        string                      `json:"accountNumber"`
	BankCode             string                      `json:"bankCode"`
	CustomerName         string                      `json:"customerName"`
	DueDate              time.Time                   `json:"dueDate"`
	RefundStatus         PaymentRefundStatus         `json:"refundStatus"`
	Expired              bool                        `json:"expired"`
	SettlementStatus     PaymentSettlementStatus     `json:"settlementStatus"`
	RefundReceiveAccount PaymentRefundReceiveAccount `json:"refundReceiveAccount,omitempty"`
}

type PaymentMobilePhone struct {
	CustomerMobilePhone string                  `json:"customerMobilePhone"`
	SettlementStatus    PaymentSettlementStatus `json:"settlementStatus"`
	ReceiptUrl          string                  `json:"receiptUrl"`
}

type PaymentGiftCertificate struct {
	ApproveNo        string                  `json:"approveNo"`
	SettlementStatus PaymentSettlementStatus `json:"settlementStatus"`
}

type PaymentTransfer struct {
	BankCode         string                  `json:"bankCode"`
	SettlementStatus PaymentSettlementStatus `json:"settlementStatus"`
}

type PaymentEasyPay struct {
	Provider       string `json:"provider"`
	Amount         int    `json:"amount"`
	DiscountAmount int    `json:"discountAmount"`
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
	Type                PaymentType             `json:"type"`
	OrderId             string                  `json:"orderId"`
	OrderName           string                  `json:"orderName"`
	MID                 string                  `json:"mId"`
	Currency            string                  `json:"currency"`
	Method              Method                  `json:"method,omitempty"`
	TotalAmount         int                     `json:"totalAmount"`
	BalanceAmount       int                     `json:"balanceAmount"`
	Status              Status                  `json:"status"`
	RequestedAt         time.Time               `json:"requestedAt"`
	ApprovedAt          *time.Time              `json:"approvedAt,omitempty"`
	UseEscrow           bool                    `json:"useEscrow"`
	LastTransactionKey  string                  `json:"lastTransactionKey,omitempty"`
	SuppliedAmount      int                     `json:"suppliedAmount"`
	Vat                 int                     `json:"vat"`
	CultureExpense      bool                    `json:"cultureExpense"`
	TaxFreeAmount       int                     `json:"taxFreeAmount"`
	TaxExemptionAmount  int                     `json:"taxExemptionAmount"`
	Cancels             []PaymentCancels        `json:"cancels,omitempty"`
	IsPartialCancelable bool                    `json:"isPartialCancelable"`
	Card                *PaymentCard            `json:"card,omitempty"`
	VirtualAccount      *PaymentVirtualAccount  `json:"virtualAccount,omitempty"`
	Secret              string                  `json:"secret,omitempty"`
	MobilePhone         *PaymentMobilePhone     `json:"mobilePhone,omitempty"`
	GiftCertificate     *PaymentGiftCertificate `json:"giftCertificate,omitempty"`
	Transfer            *PaymentTransfer        `json:"transfer,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Receipt             *PaymentURL             `json:"receipt,omitempty"`
	Checkout            *PaymentURL             `json:"checkout,omitempty"`
	EasyPay             *PaymentEasyPay         `json:"easyPay,omitempty"`
	Country             string                  `json:"country"`
	Failure             *PaymentFailure         `json:"failure,omitempty"`
	CashReceipt         *PaymentCashReceipt     `json:"cashReceipt,omitempty"`
	CashReceipts        []PaymentCashReceipts   `json:"cashReceipts,omitempty"`
	Discount            *PaymentDiscount        `json:"discount,omitempty"`
}
