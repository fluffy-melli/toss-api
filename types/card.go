package types

type Card string

const (
	CardCredit  Card = "신용"
	CardDebit   Card = "체크"
	CardGift    Card = "기프트"
	CardUnknown Card = "미확인"
)
