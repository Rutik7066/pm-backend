package createaccount

import (
	"backend/db"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateAccount(c *fiber.Ctx) error {
	// binding
	var req ConfirmationRequest
	if parError := c.BodyParser(&req); parError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Param",
			"error":   parError.Error(),
		})
	}
	data := req.Razorpay_order_id + "|" + req.Razorpay_payment_id
	fmt.Printf("Secret: %s Data: %s\n", "vAr2uK7Gix5vp0I2ugJkxrWX", data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
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

		return c.SendStatus(fiber.StatusPaymentRequired)
	}
	//  action according to
	// if payment.OrderStatus == "PAID" && req.PlanPrice == float64(payment.OrderAmount) {
	customer := GetCustomerFromReq(&req)
	result := db.AddUser(&customer)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messege": "Internal Error contact us",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&customer)
	// } else {
	// 	return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
	// 		"message": "Payment" + payment.OrderStatus,
	// 	})
	// }

}
