package bankcode

type Bank string

const (
	// 은행 코드 (숫자)

	KDBBANK_BANK     Bank = "02" // KDB산업은행
	IBK_BANK         Bank = "03" // IBK기업은행
	KOOKMIN_BANK     Bank = "04" // KB국민은행
	SUHYEOP_BANK     Bank = "07" // Sh수협은행
	NONGHYEOP_BANK   Bank = "11" // NH농협은행
	LOCALNONGHYEOP   Bank = "12" // 단위농협(지역농축협)
	WOORI_BANK       Bank = "20" // 우리은행
	SC_BANK          Bank = "23" // SC제일은행
	CITI_BANK        Bank = "27" // 씨티은행
	DAEGUBANK        Bank = "31" // DGB대구은행
	BUSANBANK        Bank = "32" // 부산은행
	GWANGJUBANK_BANK Bank = "34" // 광주은행
	JEJUBANK_BANK    Bank = "35" // 제주은행
	JEONBUKBANK_BANK Bank = "37" // 전북은행
	KYONGNAMBANK     Bank = "39" // 경남은행
	SAEMAUL_BANK     Bank = "45" // 새마을금고
	SHINHYEOP_BANK   Bank = "48" // 신협
	SAVINGBANK_BANK  Bank = "50" // 저축은행중앙회
	HSBC_BANK        Bank = "54" // 홍콩상하이은행
	SANLIM           Bank = "64" // 산림조합
	POST_BANK        Bank = "71" // 우체국예금보험
	HANA_BANK        Bank = "81" // 하나은행
	SHINHAN_BANK     Bank = "88" // 신한은행
	KBANK_BANK       Bank = "89" // 케이뱅크
	KAKAOBANK_BANK   Bank = "90" // 카카오뱅크
	TOSSBANK_BANK    Bank = "92" // 토스뱅크
)

func (b Bank) String() string {
	switch b {
	case KDBBANK_BANK:
		return "KDB산업은행"
	case IBK_BANK:
		return "IBK기업은행"
	case KOOKMIN_BANK:
		return "KB국민은행"
	case SUHYEOP_BANK:
		return "Sh수협은행"
	case NONGHYEOP_BANK:
		return "NH농협은행"
	case LOCALNONGHYEOP:
		return "단위농협(지역농축협)"
	case WOORI_BANK:
		return "우리은행"
	case SC_BANK:
		return "SC제일은행"
	case CITI_BANK:
		return "씨티은행"
	case DAEGUBANK:
		return "DGB대구은행"
	case BUSANBANK:
		return "부산은행"
	case GWANGJUBANK_BANK:
		return "광주은행"
	case JEJUBANK_BANK:
		return "제주은행"
	case JEONBUKBANK_BANK:
		return "전북은행"
	case KYONGNAMBANK:
		return "경남은행"
	case SAEMAUL_BANK:
		return "새마을금고"
	case SHINHYEOP_BANK:
		return "신협"
	case SAVINGBANK_BANK:
		return "저축은행중앙회"
	case HSBC_BANK:
		return "홍콩상하이은행"
	case SANLIM:
		return "산림조합"
	case POST_BANK:
		return "우체국예금보험"
	case HANA_BANK:
		return "하나은행"
	case SHINHAN_BANK:
		return "신한은행"
	case KBANK_BANK:
		return "케이뱅크"
	case KAKAOBANK_BANK:
		return "카카오뱅크"
	case TOSSBANK_BANK:
		return "토스뱅크"
	default:
		return "알 수 없는 은행"
	}
}
