package types

type InterestPayer string

const (
	InterestPayerBuyer       InterestPayer = "BUYER"        // 구매자
	InterestPayerCardCompany InterestPayer = "CARD_COMPANY" // 카드사
	InterestPayerMerchant    InterestPayer = "MERCHANT"     // 상점
)
