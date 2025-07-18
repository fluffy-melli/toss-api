package toss

import "time"

type Billing struct {
	MID             string      `json:"mId"`
	BillingKey      string      `json:"billingKey"`
	CustomerKey     string      `json:"customerKey"`
	AuthenticatedAt time.Time   `json:"authenticatedAt"`
	Method          Method      `json:"method"`
	Card            PaymentCard `json:"card"`
	CardCompany     string      `json:"cardCompany"`
	CardNumber      string      `json:"cardNumber"`
}
