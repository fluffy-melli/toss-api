package status

type Settlement string

const (
	SettlementIncompleted Settlement = "INCOMPLETED" // 처리중
	SettlementCompleted   Settlement = "COMPLETED"   // 처리완료
)
