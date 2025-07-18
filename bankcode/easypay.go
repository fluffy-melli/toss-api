package bankcode

type EasyPay string

const (
	// 간편결제사 코드
	TOSSPAY    EasyPay = "TP" // 토스페이
	NAVERPAY   EasyPay = "NP" // 네이버페이
	SAMSUNGPAY EasyPay = "SP" // 삼성페이
	APPLEPAY   EasyPay = "AP" // 애플페이
	LPAY       EasyPay = "LP" // 엘페이
	KAKAOPAY   EasyPay = "KP" // 카카오페이
	PAYCO      EasyPay = "PC" // 페이코
	SSG        EasyPay = "SG" // SSG페이
)
