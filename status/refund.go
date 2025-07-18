package status

type Refund string

const (
	RefundNone          Refund = "NONE"
	RefundPending       Refund = "PENDING"
	RefundFailed        Refund = "FAILED"
	RefundPartialFailed Refund = "PARTIAL_FAILED"
	RefundCompleted     Refund = "COMPLETED"
)
