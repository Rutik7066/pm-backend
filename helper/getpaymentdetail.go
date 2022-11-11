package helper

import (
	"backend/modal"

	"github.com/goccy/go-json"

	"io"
	"net/http"
)

func GetPaymentDetail(OrderId string) (detail modal.PaymentDetail, err error) {

	url := "https://api.cashfree.com/pg/orders/" + OrderId
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		err = reqErr
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-client-id", "188439f1fea02d030392616e32934881")
	req.Header.Add("x-client-secret", "4c5a9de49fe5ed5a60a7db12f14f16f0790a6d9c")
	req.Header.Add("x-api-version", "2022-01-01")
	res, reqError := http.DefaultClient.Do(req)
	if reqError != nil {
		err = reqErr
		return
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		err = readErr
		return
	}
	var paymentDetails modal.PaymentDetail
	if unErr := json.Unmarshal(body, &paymentDetails); unErr != nil {
		err = unErr
	}
	detail = paymentDetails
	return
}
