package toss

import "time"

type SettlementFeesType string

const (
	SettlementFeesTypeBase                SettlementFeesType = "BASE"
	SettlementFeesTypeInstallment         SettlementFeesType = "INSTALLMENT"
	SettlementFeesTypeInstallmentDiscount SettlementFeesType = "INSTALLMENT_DISCOUNT"
	SettlementFeesTypePointSaving         SettlementFeesType = "POINT_SAVING"
	SettlementFeesTypeEtc                 SettlementFeesType = "ETC"
)

type SettlementFee struct {
	Type SettlementFeesType `json:"type"`
	Fee  int                `json:"fee"`
}

type Settlement struct {
	MID             string                  `json:"mId"`
	PaymentKey      string                  `json:"paymentKey"`
	TransactionKey  string                  `json:"transactionKey"`
	OrderId         string                  `json:"orderId"`
	Currency        string                  `json:"currency"`
	Method          Method                  `json:"method"`
	Amount          int                     `json:"amount"`
	InterestFee     int                     `json:"interestFee"`
	Fees            []SettlementFee         `json:"fees"`
	SupplyAmount    int                     `json:"supplyAmount"`
	Vat             int                     `json:"vat"`
	PayOutAmount    int                     `json:"payOutAmount"`
	ApprovedAt      time.Time               `json:"approvedAt"`
	SoldDate        string                  `json:"soldDate"`
	PaidOutDate     string                  `json:"paidOutDate"`
	Card            *PaymentCard            `json:"card,omitempty"`
	EasyPay         *PaymentEasyPay         `json:"easyPay,omitempty"`
	GiftCertificate *PaymentGiftCertificate `json:"giftCertificate,omitempty"`
	MobilePhone     *PaymentMobilePhone     `json:"mobilePhone,omitempty"`
	Transfer        *PaymentTransfer        `json:"transfer,omitempty"`
	VirtualAccount  *PaymentVirtualAccount  `json:"virtualAccount,omitempty"`
	Cancels         []PaymentCancels        `json:"cancels,omitempty"`
}
