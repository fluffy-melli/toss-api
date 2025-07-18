package types

type Payment string

const (
	PaymentNormal   Payment = "NORMAL"   // 일반결제
	PaymentBilling  Payment = "BILLING"  // 자동결제
	PaymentBrandpay Payment = "BRANDPAY" // 브랜드페이
)
