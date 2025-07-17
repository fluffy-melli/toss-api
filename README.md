
> [!WARNING]
> 해당 모듈은 **Toss Payments 비공식 모듈** 입니다.

Toss Payments API를 위한 Go 클라이언트 입니다. 결제 승인, 취소, 조회 등의 핵심 기능을 제공합니다.

```sh
go get -u github.com/fluffy-melli/toss-api
```

| 코어 API 항목       | 구현 여부 | 테스트용 키 | 실사용 키 |
|----------------|----|----|----|
| 결제            | O | O | ? (테스트 안함) |
| 자동결제        | X | X (미구현) | X (미구현) |
| 거래내역         | O | O | ? (테스트 안함) |
| 정산내역         | O | X (응답은 빈배열) | ? (테스트 안함) |
| 현금영수증      | X | X (미구현) | X (미구현) |
| 지급대행        | X | X (미구현) | X (미구현) |
| 프로모션        | X | X (미구현) | X (미구현) |
| 프렌드페이 API  | X | X (미구현) | X (미구현) |

자세한 API 문서는 [Toss Payments 공식 API 레퍼런스](https://docs.tosspayments.com/reference)를 참고하세요.