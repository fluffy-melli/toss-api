
> [!WARNING]
> 이 프로젝트는 **Toss Payments 비공식 라이브러리** 입니다.

Toss Payments API를 위한 Go 클라이언트 입니다. 결제 승인, 취소, 조회 등의 핵심 기능을 제공합니다.

```sh
go get -u github.com/fluffy-melli/toss-api
```

| 코어 API 항목| 구현 여부 | 테스트키 | 실사용키 |
|----------------|----|----|----|
| 일반결제 | O | O | ? (미테스트) |
| 자동결제 | O | O | ? (미테스트) |
| 거래내역 | O | O | ? (미테스트) |
| 정산내역 | O | X (빈배열 응답) | ? (미테스트) |
| 지급대행 | X | X (미구현) | X (미구현) |
| 프로모션 | X | X (미구현) | X (미구현) |
| 현금영수증 | X | X (미구현) | X (미구현) |
| 프렌드페이 | X | X (미구현) | X (미구현) |

자세한 API 문서는 [Toss Payments 공식 API 레퍼런스](https://docs.tosspayments.com/reference)를 참고하세요.