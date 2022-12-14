package credit

import (
	"backend/db"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdateCredit(c *fiber.Ctx) error {
	type creditReq struct {
		Razorpay_payment_id string `json:"razorpay_payment_id"`
		Razorpay_order_id   string `json:"razorpay_order_id"`
		Razorpay_signature  string `json:"razorpay_signature"`
		Uid                 string `json:"uid"`
		PlanName            string `json:"planname"`
	}
	var req creditReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messege": "Invalid Param",
		})
	}
	data := req.Razorpay_order_id + "|" + req.Razorpay_payment_id
	fmt.Printf("Secret: %s Data: %s\n", "vAr2uK7Gix5vp0I2ugJkxrWX", data)
	h := hmac.New(sha256.New, []byte("vAr2uK7Gix5vp0I2ugJkxrWX"))

	// Write Data to it
	_, err := h.Write([]byte(data))

	if err != nil {
		panic(err)
	}

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("Result: %s\n", sha)

	if subtle.ConstantTimeCompare([]byte(sha), []byte(req.Razorpay_signature)) != 1 {
		fmt.Println("Works")
		return c.SendStatus(fiber.StatusPaymentRequired)

	}
	var credit int
	switch req.PlanName {
	case "250":
		credit = 250
	case "500":
		credit = 500
	case "1000":
		credit = 1500
	case "3000":
		credit = 5000
	case "5000":
		credit = 10000
	}
	// amt, _ := strconv.Atoi(req.PlanName)
	// if paymentDetail.OrderStatus == "PAID" && paymentDetail.OrderAmount == amt {
	user := db.AddCredit(req.Uid, credit)
	fmt.Println(user)
	return c.Status(fiber.StatusOK).JSON(&user)
	// } else {
	// 	return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
	// 		"messege": "Payment" + paymentDetail.OrderStatus,
	// 	})
	// }

}
