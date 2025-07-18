package toss

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	API = "https://api.tosspayments.com/"
)

type Client struct {
	apiKey        string
	encodedApiKey string
}

/*
NewClient 토스페이먼츠 클라이언트 생성
토스페이먼츠 API를 사용하기 위한 클라이언트를 생성합니다.
API 키는 토스페이먼츠 개발자 센터에서 발급받을 수 있습니다.

Parameters:
  - apiKey: 토스페이먼츠에서 발급받은 API 키 (시크릿 키)

Returns:
  - *Client: 초기화된 토스페이먼츠 클라이언트 객체

Example:

	client := toss.NewClient("test_sk_zXLkKEypNArWmo50nX3lmeaxYG5R")
	payment, err := client.PaymentConfirm("paymentKey", "orderId", 15000)
*/
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:        apiKey,
		encodedApiKey: base64.StdEncoding.EncodeToString([]byte(apiKey + ":")),
	}
}

func (c *Client) req(method, urls string, body map[string]any) ([]byte, error) {
	var req *http.Request
	var err error

	fullURL := API + urls

	if method == "GET" || body == nil {
		params := url.Values{}
		for k, v := range body {
			params.Set(k, fmt.Sprintf("%v", v))
		}
		fullURL += "?" + params.Encode()
		fmt.Println(fullURL)
		req, err = http.NewRequest(method, fullURL, nil)
	} else {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, fullURL, bytes.NewBuffer(jsonData))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.encodedApiKey))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[%d] Toss API error: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *Client) paymentReq(method, urls string, body map[string]any) (*Payment, error) {
	respBody, err := c.req(method, urls, body)

	if err != nil {
		return nil, err
	}

	var payment Payment

	if err := json.Unmarshal(respBody, &payment); err != nil {
		return nil, err
	}

	return &payment, nil
}

func (c *Client) transactionsReq(method, urls string, body map[string]any) ([]Transaction, error) {
	respBody, err := c.req(method, urls, body)

	if err != nil {
		return nil, err
	}

	var transactions []Transaction

	if err := json.Unmarshal(respBody, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *Client) settlementsReq(method, urls string, body map[string]any) ([]Settlement, error) {
	respBody, err := c.req(method, urls, body)

	if err != nil {
		return nil, err
	}

	var settlements []Settlement

	if err := json.Unmarshal(respBody, &settlements); err != nil {
		return nil, err
	}

	return settlements, nil
}

func (c *Client) billingReq(method, urls string, body map[string]any) (*Billing, error) {
	respBody, err := c.req(method, urls, body)

	if err != nil {
		return nil, err
	}

	var billing Billing

	if err := json.Unmarshal(respBody, &billing); err != nil {
		return nil, err
	}

	return &billing, nil
}

/*
PaymentConfirm 결제 승인
토스페이먼츠 결제 위젯에서 결제 인증 완료 후 반드시 호출해야 하는 API입니다.
결제 인증 후 10분 이내에 호출하지 않으면 결제가 자동으로 만료됩니다.

Parameters:
  - paymentKey: 결제 위젯에서 전달받은 결제 고유 키
  - orderId: 결제 요청 시 생성한 주문 ID (최대 64자)
  - amount: 실제 결제 금액 (위젯에서 전달받은 금액과 일치해야 함)

Returns:
  - *Payment: 승인된 결제 정보 객체
  - error: 승인 실패 시 에러 (금액 불일치, 만료, 중복 승인 등)

Example:

	payment, err := client.PaymentConfirm("tgen_20240101000000_ABC123", "order_12345", 15000)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("결제 승인 완료: %s\n", payment.OrderID)
*/
func (c *Client) PaymentConfirm(paymentKey, orderId string, amount int) (*Payment, error) {
	return c.paymentReq("POST", "v1/payments/confirm", map[string]any{
		"paymentKey": paymentKey,
		"orderId":    orderId,
		"amount":     amount,
	})
}

/*
PaymentCancel 결제 취소
승인된 결제를 전액 또는 부분 취소합니다.
취소 후에는 되돌릴 수 없으니 신중하게 사용하세요.

Parameters:
  - paymentKey: 취소할 결제의 고유 키
  - cancelReason: 취소 사유 (필수, 최대 200자)
  - o: 취소 옵션 (nil 가능)
  - CancelAmount: 부분 취소 금액 (nil이면 전액 취소)
  - RefundReceiveAccount: 환불 계좌 정보 (계좌이체 결제시 필요)
  - TaxFreeAmount: 면세 금액
  - Currency: 통화 코드

Returns:
  - *Payment: 취소 처리된 결제 정보 객체
  - error: 취소 실패 시 에러 (이미 취소됨, 취소 불가 상태 등)

Example:

	// 전액 취소
	payment, err := client.PaymentCancel("tgen_20240101000000_ABC123", "고객 요청", nil)

	// 부분 취소 (5,000원만 취소)
	options := &PaymentCancelOptions{
	    CancelAmount: 5000,
	}
	payment, err := client.PaymentCancel("tgen_20240101000000_ABC123", "부분 환불", options)
*/
func (c *Client) PaymentCancel(paymentKey, cancelReason string, o *PaymentCancelOptions) (*Payment, error) {
	body := map[string]any{}
	if o != nil {
		body["cancelAmount"] = o.CancelAmount
		body["refundReceiveAccount"] = map[string]any{
			"bank":          o.RefundReceiveAccount.Bank,
			"accountNumber": o.RefundReceiveAccount.AccountNumber,
			"holderName":    o.RefundReceiveAccount.HolderName,
		}
		body["taxFreeAmount"] = o.TaxFreeAmount
		body["currency"] = o.Currency
	}
	body["cancelReason"] = cancelReason
	return c.paymentReq("POST", fmt.Sprintf("v1/payments/%s/cancel", paymentKey), body)
}

/*
PaymentCheckByOrderID 주문 ID로 결제 조회
상점에서 생성한 주문 ID로 결제 정보를 조회합니다.
주문 상태 확인이나 결제 정보 조회 시 사용하세요.

Parameters:
  - orderId: 결제 요청 시 생성한 주문 ID

Returns:
  - *Payment: 결제 정보 객체
  - error: 조회 실패 시 에러 (존재하지 않는 주문 등)

Example:

	payment, err := client.PaymentCheckByOrderID("order_12345")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("결제 상태: %s, 금액: %d원\n", payment.Status, payment.TotalAmount)
*/
func (c *Client) PaymentCheckByOrderID(orderId string) (*Payment, error) {
	return c.paymentReq("GET", fmt.Sprintf("v1/payments/orders/%s", orderId), nil)
}

/*
PaymentCheckByPaymentKey 결제 키로 결제 조회
토스페이먼츠에서 발급한 결제 키로 결제 정보를 조회합니다.
결제 승인 후 받은 paymentKey로 언제든지 결제 정보를 확인할 수 있습니다.

Parameters:
  - paymentKey: 결제 승인 후 받은 결제 고유 키

Returns:
  - *Payment: 결제 정보 객체
  - error: 조회 실패 시 에러 (존재하지 않는 결제 등)

Example:

	payment, err := client.PaymentCheckByPaymentKey("tgen_20240101000000_ABC123")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("결제 방법: %s, 승인 시간: %s\n", payment.Method, payment.ApprovedAt)
*/
func (c *Client) PaymentCheckByPaymentKey(paymentKey string) (*Payment, error) {
	return c.paymentReq("GET", fmt.Sprintf("v1/payments/%s", paymentKey), nil)
}

/*
TransactionsCheck 거래 내역 조회
특정 기간의 모든 거래 기록을 조회합니다.
정산, 매출 분석, 거래 내역 확인 등에 활용하세요.

Parameters:
  - startDate: 조회 시작 날짜 (거래 처리 시점 기준)
  - endDate: 조회 종료 날짜 (거래 처리 시점 기준)
  - o: 조회 옵션 (nil 가능)
  - startingAfter: 페이징을 위한 커서 (이전 조회 결과의 마지막 거래 ID)
  - limit: 조회할 거래 수 (최대 100개, 기본값 20개)

Returns:
  - []Transaction: 거래 내역 배열
  - error: 조회 실패 시 에러

Example:

	// 최근 7일 거래 내역 조회
	startDate := time.Now().AddDate(0, 0, -7)
	endDate := time.Now()
	transactions, err := client.TransactionsCheck(startDate, endDate, nil)

	// 페이징으로 100개씩 조회
	options := &TransactionOptions{
	    Limit: 100,
	}
	transactions, err := client.TransactionsCheck(startDate, endDate, options)
*/
func (c *Client) TransactionsCheck(startDate, endDate time.Time, o *TransactionOptions) ([]Transaction, error) {
	body := map[string]any{}
	if o != nil {
		body["startingAfter"] = o.StartingAfter
		body["limit"] = o.Limit
	}
	body["startDate"] = startDate.Format(time.RFC3339)
	body["endDate"] = endDate.Format(time.RFC3339)
	return c.transactionsReq("GET", "v1/transactions", body)
}

/*
SettlementsCheck 정산 내역 조회
지정한 날짜 정보로 정산 기록을 조회합니다.
정산 관리, 매출 분석, 수수료 확인 등에 활용하세요.

Parameters:
  - startDate: 조회 시작 날짜 (정산 처리 시점 기준)
  - endDate: 조회 종료 날짜 (정산 처리 시점 기준)
  - o: 조회 옵션 (nil 가능)
  - Page: 페이지 번호 (1부터 시작)
  - Size: 페이지당 조회할 정산 수 (최대 100개, 기본값 20개)

Returns:
  - []Settlement: 정산 내역 배열
  - error: 조회 실패 시 에러

Example:

	// 최근 30일 정산 내역 조회
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()
	settlements, err := client.SettlementsCheck(startDate, endDate, nil)

	// 페이징으로 50개씩 조회
	options := &SettlementOptions{
	    Page: 1,
	    Size: 50,
	}
	settlements, err := client.SettlementsCheck(startDate, endDate, options)
*/
func (c *Client) SettlementsCheck(startDate, endDate time.Time, o *SettlementOptions) ([]Settlement, error) {
	body := map[string]any{}
	if o != nil {
		body["page"] = o.Page
		body["size"] = o.Size
	}
	body["startDate"] = startDate.Format(time.RFC3339)
	body["endDate"] = endDate.Format(time.RFC3339)
	return c.settlementsReq("GET", "v1/settlements", body)
}

/*
BillingIssue 빌링키 발급
고객의 결제 수단을 등록하고 빌링키를 발급합니다.
자동결제(정기결제) 서비스에 활용하세요.

Parameters:
  - authKey: 결제 수단 인증 키
  - customerKey: 고객 식별 키

Returns:
  - *Billing: 빌링키 정보
  - error: 발급 실패 시 에러

Example:

	// 빌링키 발급
	billing, err := client.BillingIssue("auth_key_123", "customer_001")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("빌링키: %s\n", billing.BillingKey)
*/
func (c *Client) BillingIssue(authKey, customerKey string) (*Billing, error) {
	return c.billingReq("POST", "v1/billing/authorizations/issue", map[string]any{
		"authKey":     authKey,
		"customerKey": customerKey,
	})
}

/*
BillingConfirm 빌링키 결제 승인
발급된 빌링키를 사용하여 결제를 승인합니다.
자동결제(정기결제) 처리에 활용하세요.

Parameters:
  - billingKey: 빌링키
  - customerKey: 고객 식별 키
  - orderId: 주문 ID
  - orderName: 주문명
  - amount: 결제 금액
  - o: 결제 옵션 (nil 가능)
  - CustomerEmail: 고객 이메일
  - CustomerName: 고객 이름
  - TaxFreeAmount: 비과세 금액
  - TaxExemptionAmount: 과세 면제 금액

Returns:
  - *Payment: 결제 정보
  - error: 결제 실패 시 에러

Example:

	// 기본 빌링키 결제
	payment, err := client.BillingConfirm("billing_key_123", "customer_001", "order_456", "정기결제", 10000, nil)
	// 옵션 포함 빌링키 결제
	options := &BillingOptions{
	    CustomerEmail: "customer@example.com",
	    CustomerName: "홍길동",
	    TaxFreeAmount: 1000,
	}
	payment, err := client.BillingConfirm("billing_key_123", "customer_001", "order_456", "정기결제", 10000, options)
*/
func (c *Client) BillingConfirm(billingKey, customerKey, orderId, orderName string, amount int, o *BillingOptions) (*Payment, error) {
	body := map[string]any{}
	if o != nil {
		body["customerEmail"] = o.CustomerEmail
		body["customerName"] = o.CustomerName
		body["taxFreeAmount"] = o.TaxFreeAmount
		body["taxExemptionAmount"] = o.TaxExemptionAmount
	}
	body["customerKey"] = customerKey
	body["orderId"] = orderId
	body["orderName"] = orderName
	body["amount"] = amount
	return c.paymentReq("POST", fmt.Sprintf("v1/billing/%s", billingKey), body)
}
