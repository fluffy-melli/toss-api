package types

type Fees string

const (
	FeesBase                Fees = "BASE"                 // 기본 수수료
	FeesInstallment         Fees = "INSTALLMENT"          // 상점
	FeesInstallmentDiscount Fees = "INSTALLMENT_DISCOUNT" // 카드사
	FeesPointSaving         Fees = "POINT_SAVING"         // 포인트 적립
	FeesEtc                 Fees = "ETC"                  // 기타
)
