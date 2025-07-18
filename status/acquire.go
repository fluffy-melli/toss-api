package status

type Acquire string

const (
	AcquireReady           Acquire = "READY"
	AcquireRequested       Acquire = "REQUESTED"
	AcquireCompleted       Acquire = "COMPLETED"
	AcquireCancelRequested Acquire = "CANCEL_REQUESTED"
	AcquireCanceled        Acquire = "CANCELED"
)
