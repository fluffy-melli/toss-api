package types

type CashReceipt string

const (
	CashReceiptIncomeDeduction = "소득공제"
	CashReceiptExpenseProof    = "지출증빙"
	CashReceiptNotIssued       = "미발행"
)
