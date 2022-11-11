package createaccount

import (
	"backend/modal"
	"time"
)

type ConfirmationRequest struct {
	OrderId          string  `json:"order_id"`
	CustomerName     string  `json:"customer_name"`
	CustomerPhone    string  `json:"customer_phone"`
	CustomerAltPhone string  `json:"customer_alt_phone"`
	CustomerEmail    string  `json:"customer_email"`
	CustomerPass     string  `json:"password"`
	BusinessAddress  string  `json:"business_address"`
	FbId             string  `json:"fb_id"`
	SnapId           string  `json:"snap_id"`
	InstaId          string  `json:"insta_id"`
	Web              string  `json:"web"`
	IpAddress        string  `json:"ip_address"`
	PlanPrice        float64 `json:"plan_price"`
}

func GetCustomerFromReq(reqUser *ConfirmationRequest) modal.Customer {
	return modal.Customer{
		CustomerName:     reqUser.CustomerName,
		CustomerEmail:    reqUser.CustomerEmail,
		CustomerPhone:    reqUser.CustomerPhone,
		CustomerAltPhone: reqUser.CustomerAltPhone,
		CustomerPass:     reqUser.CustomerPass,
		BusinessAddress:  reqUser.BusinessAddress,
		FbId:             reqUser.FbId,
		SnapId:           reqUser.SnapId,
		InstaId:          reqUser.InstaId,
		Web:              reqUser.Web,
		IpAddress:        reqUser.IpAddress,
		PlanPrice:        reqUser.PlanPrice,
		ValidTill:        time.Now().Add(time.Hour * 8760),
		Credit:           15,
	}
}
