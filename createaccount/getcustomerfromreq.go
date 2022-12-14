package createaccount

import (
	"backend/modal"
	"time"
)

type ConfirmationRequest struct {
	Razorpay_payment_id string  `json:"razorpay_payment_id"`
	Razorpay_order_id   string  `json:"razorpay_order_id"`
	Razorpay_signature  string  `json:"razorpay_signature"`
	CustomerName        string  `json:"customername"`
	CustomerPhone       string  `json:"customerphone"`
	CustomerAltPhone    string  `json:"customeraltphone"`
	CustomerEmail       string  `json:"customeremail"`
	CustomerPass        string  `json:"password"`
	BussinesName        string  `json:"businessname"`
	BusinessAddress     string  `json:"businessaddress"`
	PlanPrice           float64 `json:"planprice"`
}

func GetCustomerFromReq(reqUser *ConfirmationRequest) modal.Customer {
	return modal.Customer{
		CustomerName:     reqUser.CustomerName,
		CustomerEmail:    reqUser.CustomerEmail,
		CustomerPhone:    reqUser.CustomerPhone,
		CustomerAltPhone: reqUser.CustomerAltPhone,
		CustomerPass:     reqUser.CustomerPass,
		BusinessAddress:  reqUser.BusinessAddress,
		BusinessName:     reqUser.BussinesName,
		PlanPrice:        reqUser.PlanPrice,
		ValidTill:        time.Now().Add(time.Hour * 8760),
		Credit:           15,
	}
}
