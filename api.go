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

func (c *Client) PaymentConfirm(paymentKey, orderId string, amount int) (*Payment, error) {
	return c.paymentReq("POST", "v1/payments/confirm", map[string]any{
		"paymentKey": paymentKey,
		"orderId":    orderId,
		"amount":     amount,
	})
}

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

func (c *Client) PaymentCheckByOrderID(orderId string) (*Payment, error) {
	return c.paymentReq("GET", fmt.Sprintf("v1/payments/orders/%s", orderId), nil)
}

func (c *Client) PaymentCheckByPaymentKey(paymentKey string) (*Payment, error) {
	return c.paymentReq("GET", fmt.Sprintf("v1/payments/%s", paymentKey), nil)
}

func (c *Client) TransactionsCheck(startDate, endDate time.Time, o *TransactionsOptions) ([]Transaction, error) {
	body := map[string]any{}
	if o != nil {
		body["startingAfter"] = o.startingAfter
		body["limit"] = o.limit
	}
	body["startDate"] = startDate.Format(time.RFC3339)
	body["endDate"] = endDate.Format(time.RFC3339)
	return c.transactionsReq("GET", "v1/transactions", body)
}
