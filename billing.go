package toss

import (
	"time"

	"github.com/fluffy-melli/toss-api/types"
)

type Billing struct {
	MID             string       `json:"mId"`
	BillingKey      string       `json:"billingKey"`
	CustomerKey     string       `json:"customerKey"`
	AuthenticatedAt time.Time    `json:"authenticatedAt"`
	Method          types.Method `json:"method"`
	Card            PaymentCard  `json:"card"`
	CardCompany     string       `json:"cardCompany"`
	CardNumber      string       `json:"cardNumber"`
}
