package gettoken

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetToken(c *fiber.Ctx) error {
	type TokenReq struct {
		Amount        string `json:"order_amount"`
		Note          string `json:"order_note"`
		CustomerUid   string    `json:"uid"`
		CustomerName  string `json:"customer_name"`
		CustomerEmail string `json:"customer_email"`
		CustomerPhone string `json:"customer_phone"`
	}
	var tokenReq TokenReq
	if err := c.BodyParser(&tokenReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "InValid Param",
			"error":   err.Error(),
		})
	}
	data, err := CreatToken(tokenReq.Amount, tokenReq.Note, tokenReq.CustomerName, tokenReq.CustomerEmail, tokenReq.CustomerPhone,tokenReq.CustomerUid)
	fmt.Println(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&data)

}
