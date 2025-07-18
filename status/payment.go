package status

type Payment string

const (
	PaymentReady             Payment = "READY"
	PaymentInProgress        Payment = "IN_PROGRESS"
	PaymentWaitingForDeposit Payment = "WAITING_FOR_DEPOSIT"
	PaymentDone              Payment = "DONE"
	PaymentCanceled          Payment = "CANCELED"
	PaymentPartialCanceled   Payment = "PARTIAL_CANCELED"
	PaymentAborted           Payment = "ABORTED"
	PaymentExpired           Payment = "EXPIRED"
)
