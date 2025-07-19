package bankcode

type Card string

const (
	// 국내 카드사 코드

	IBK_BC      Card = "3K" // 기업 BC
	GWANGJUBANK Card = "46" // 광주은행
	LOTTE       Card = "71" // 롯데카드
	KDBBANK     Card = "30" // KDB산업은행
	BC          Card = "31" // BC카드
	SAMSUNG     Card = "51" // 삼성카드
	SAEMAUL     Card = "38" // 새마을금고
	SHINHAN     Card = "41" // 신한카드
	SHINHYEOP   Card = "62" // 신협
	CITI        Card = "36" // 씨티카드
	WOORI_BC    Card = "33" // 우리BC카드(BC 매입)
	WOORI       Card = "W1" // 우리카드(우리 매입)
	POST        Card = "37" // 우체국예금보험
	SAVINGBANK  Card = "39" // 저축은행중앙회
	JEONBUKBANK Card = "35" // 전북은행
	JEJUBANK    Card = "42" // 제주은행
	KAKAOBANK   Card = "15" // 카카오뱅크
	KBANK       Card = "3A" // 케이뱅크
	TOSSBANK    Card = "24" // 토스뱅크
	HANA        Card = "21" // 하나카드
	HYUNDAI     Card = "61" // 현대카드
	KOOKMIN     Card = "11" // KB국민카드
	NONGHYEOP   Card = "91" // NH농협카드
	SUHYEOP     Card = "34" // Sh수협은행

	// 해외 카드사 코드

	DINERS   Card = "6D" // 다이너스 클럽
	MASTER   Card = "4M" // 마스터카드
	UNIONPAY Card = "3C" // 유니온페이
	AMEX     Card = "7A" // 아메리칸 익스프레스
	JCB      Card = "4J" // JCB
	VISA     Card = "4V" // VISA

	// 증권사 코드

	KYOBO_SECURITIES                 Card = "261" // 교보증권
	DAISHIN_SECURITIES               Card = "267" // 대신증권
	MERITZ_SECURITIES                Card = "287" // 메리츠증권
	MIRAE_ASSET_SECURITIES           Card = "238" // 미래에셋증권
	BOOKOOK_SECURITIES               Card = "290" // 부국증권
	SAMSUNG_SECURITIES               Card = "240" // 삼성증권
	SHINYOUNG_SECURITIES             Card = "291" // 신영증권
	SHINHAN_INVESTMENT               Card = "278" // 신한금융투자
	YUANTA_SECURITES                 Card = "209" // 유안타증권
	EUGENE_INVESTMENT_AND_SECURITIES Card = "280" // 유진투자증권
	KAKAOPAY_SECURITIES              Card = "288" // 카카오페이증권
	KIWOOM                           Card = "264" // 키움증권
	TOSS_MONEY                       Card = "-"   // 토스머니
	TOSS_SECURITIES                  Card = "271" // 토스증권
	KOREA_FOSS_SECURITIES            Card = "294" // 펀드온라인코리아(한국포스증권)
	HANA_INVESTMENT_AND_SECURITIES   Card = "270" // 하나금융투자
	HI_INVESTMENT_AND_SECURITIES     Card = "262" // 하이투자증권
	KOREA_INVESTMENT_AND_SECURITIES  Card = "243" // 한국투자증권
	HANHWA_INVESTMENT_AND_SECURITIES Card = "269" // 한화투자증권
	HYUNDAI_MOTOR_SECURITIES         Card = "263" // 현대차증권
	HSBC                             Card = "054" // 홍콩상하이은행
	DB_INVESTMENT_AND_SECURITIES     Card = "279" // DB금융투자
	IBK                              Card = "003" // IBK기업은행
	KB_SECURITIES                    Card = "218" // KB증권
	DAOL_INVESTMENT_AND_SECURITIES   Card = "227" // KTB투자증권(다올투자증권)
	LIG_INVESTMENT_AND_SECURITIES    Card = "292" // LIG투자증권
	NH_INVESTMENT_AND_SECURITIES     Card = "247" // NH투자증권
	SC                               Card = "023" // SC제일은행
	SK_SECURITIES                    Card = "266" // SK증권
)

func (c Card) String() string {
	switch c {
	// 국내 카드사 코드
	case IBK_BC:
		return "기업 BC"
	case GWANGJUBANK:
		return "광주은행"
	case LOTTE:
		return "롯데카드"
	case KDBBANK:
		return "KDB산업은행"
	case BC:
		return "BC카드"
	case SAMSUNG:
		return "삼성카드"
	case SAEMAUL:
		return "새마을금고"
	case SHINHAN:
		return "신한카드"
	case SHINHYEOP:
		return "신협"
	case CITI:
		return "씨티카드"
	case WOORI_BC:
		return "우리BC카드"
	case WOORI:
		return "우리카드"
	case POST:
		return "우체국예금보험"
	case SAVINGBANK:
		return "저축은행중앙회"
	case JEONBUKBANK:
		return "전북은행"
	case JEJUBANK:
		return "제주은행"
	case KAKAOBANK:
		return "카카오뱅크"
	case KBANK:
		return "케이뱅크"
	case TOSSBANK:
		return "토스뱅크"
	case HANA:
		return "하나카드"
	case HYUNDAI:
		return "현대카드"
	case KOOKMIN:
		return "KB국민카드"
	case NONGHYEOP:
		return "NH농협카드"
	case SUHYEOP:
		return "Sh수협은행"

	// 해외 카드사 코드
	case DINERS:
		return "다이너스 클럽"
	case MASTER:
		return "마스터카드"
	case UNIONPAY:
		return "유니온페이"
	case AMEX:
		return "아메리칸 익스프레스"
	case JCB:
		return "JCB"
	case VISA:
		return "VISA"

	// 은행 & 증권사 코드
	case KYOBO_SECURITIES:
		return "교보증권"
	case DAISHIN_SECURITIES:
		return "대신증권"
	case MERITZ_SECURITIES:
		return "메리츠증권"
	case MIRAE_ASSET_SECURITIES:
		return "미래에셋증권"
	case BOOKOOK_SECURITIES:
		return "부국증권"
	case SAMSUNG_SECURITIES:
		return "삼성증권"

	case SHINYOUNG_SECURITIES:
		return "신영증권"
	case SHINHAN_INVESTMENT:
		return "신한금융투자"
	case YUANTA_SECURITES:
		return "유안타증권"
	case EUGENE_INVESTMENT_AND_SECURITIES:
		return "유진투자증권"
	case KAKAOPAY_SECURITIES:
		return "카카오페이증권"
	case KIWOOM:
		return "키움증권"
	case TOSS_MONEY:
		return "토스머니"
	case TOSS_SECURITIES:
		return "토스증권"
	case KOREA_FOSS_SECURITIES:
		return "한국포스증권"
	case HANA_INVESTMENT_AND_SECURITIES:
		return "하나금융투자"
	case HI_INVESTMENT_AND_SECURITIES:
		return "하이투자증권"
	case KOREA_INVESTMENT_AND_SECURITIES:
		return "한국투자증권"
	case HANHWA_INVESTMENT_AND_SECURITIES:
		return "한화투자증권"
	case HYUNDAI_MOTOR_SECURITIES:
		return "현대차증권"
	case HSBC:
		return "홍콩상하이은행"
	case DB_INVESTMENT_AND_SECURITIES:
		return "DB금융투자"
	case IBK:
		return "IBK기업은행"
	case KB_SECURITIES:
		return "KB증권"
	case DAOL_INVESTMENT_AND_SECURITIES:
		return "KTB투자증권"
	case LIG_INVESTMENT_AND_SECURITIES:
		return "LIG투자증권"
	case NH_INVESTMENT_AND_SECURITIES:
		return "NH투자증권"
	case SC:
		return "SC제일은행"
	case SK_SECURITIES:
		return "SK증권"
	default:
		return "알 수 없는 금융기관"
	}
}
