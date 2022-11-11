package gettoken

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func CreatToken(amount string, note string, customerName string, customerEmail string, customerPhone string,customerUid string) (data map[string]interface{}, err error) {
	request := map[string]any{
		"order_amount":   amount,
		"order_currency": "INR",
		"order_note":     note,
		"customer_details": map[string]string{
			"customer_id":    customerUid,
			"customer_name":  customerName,
			"customer_email": customerEmail,
			"customer_phone": customerPhone,
		},
	}
	jsonObj, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		err = jsonErr
		return
	}
	url := "https://api.cashfree.com/pg/orders"
	req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(jsonObj))
	if reqErr != nil {
		err = reqErr
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-version", "2022-01-01")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-client-id", os.Getenv("APPID"))
	req.Header.Add("x-client-secret", os.Getenv("SECRETEKEY"))
	res, recErr := http.DefaultClient.Do(req)
	if recErr != nil {
		err = reqErr
		return
	}
	defer res.Body.Close()
	reaData, rerr := io.ReadAll(res.Body)
	if rerr != nil {
		err = rerr
		return
	}
	var mp map[string]interface{}
	json.Unmarshal([]byte(reaData), &mp)
	if jsonErr != nil {
		err = jsonErr
		return
	}
	data = mp
	return

}
