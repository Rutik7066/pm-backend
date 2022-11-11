package modal

import "time"

type 
PaymentDetail struct {
	CfOrderID       int       `json:"cf_order_id"`
	CreatedAt       time.Time `json:"created_at"`
	CustomerDetails struct {
		CustomerID    string `json:"customer_id"`
		CustomerName  string `json:"customer_name"`
		CustomerEmail string `json:"customer_email"`
		CustomerPhone string `json:"customer_phone"`
	} `json:"customer_details"`
	Entity          string    `json:"entity"`
	OrderAmount     int       `json:"order_amount"`
	OrderCurrency   string    `json:"order_currency"`
	OrderExpiryTime time.Time `json:"order_expiry_time"`
	OrderID         string    `json:"order_id"`
	OrderMeta       struct {
		ReturnURL      interface{} `json:"return_url"`
		NotifyURL      interface{} `json:"notify_url"`
		PaymentMethods interface{} `json:"payment_methods"`
	} `json:"order_meta"`
	OrderNote   string        `json:"order_note"`
	OrderSplits []interface{} `json:"order_splits"`
	OrderStatus string        `json:"order_status"`
	OrderTags   interface{}   `json:"order_tags"`
	OrderToken  string        `json:"order_token"`
	PaymentLink string        `json:"payment_link"`
	Payments    struct {
		URL string `json:"url"`
	} `json:"payments"`
	Refunds struct {
		URL string `json:"url"`
	} `json:"refunds"`
	Settlements struct {
		URL string `json:"url"`
	} `json:"settlements"`
}
