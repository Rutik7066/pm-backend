package rpay

import (
	"log"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	razorpay "github.com/razorpay/razorpay-go"
)

func CreateOrder(c *fiber.Ctx) error {
	type req struct {
		Amount string                 `json:"amount"`
		Note   map[string]interface{} `json:"notes"`
	}
	log.Println("[REQUEST BODY] : " + string(c.Body()))
	var r req
	err := json.Unmarshal(c.Body(), &r)
	if err != nil {
		log.Println("[ERROR] : " + err.Error())
		return c.SendString(err.Error())
	}
	client := razorpay.NewClient("rzp_live_ke2XNPaoJ3IbuK", "vAr2uK7Gix5vp0I2ugJkxrWX")
	amount, err := strconv.Atoi(r.Amount)
	if err != nil {
		log.Println("[ERROR] : " + err.Error())
		return c.SendString(err.Error())
	}
	data := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"notes":    r.Note,
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Println("[ERROR] : " + err.Error())
		return c.SendString(err.Error())
	}
	// log.Println("[Request Response] : " + string(body))
	return c.Status(fiber.StatusOK).JSON(&body)

}
