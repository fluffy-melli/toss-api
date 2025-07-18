package types

type Method string

const (
	MethodCard           Method = "카드"
	MethodVirtualAccount Method = "가상계좌"
	MethodEasyPay        Method = "간편결제"
	MethodMobilePhone    Method = "핸드폰"
	MethodTransfer       Method = "계좌이체"
	MethodCulture        Method = "문화상품권"
	MethodBookCulture    Method = "도서문화상품권"
	MethodGameCulture    Method = "게임문화상품권"
)
