package createaccount

import (
	"backend/db"
	"backend/helper"
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
	// payment detail check
	payment, payErr := helper.GetPaymentDetail(req.OrderId)
	if payErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Payment Error",
			"error":   payErr.Error(),
		})
	}
	fmt.Println(payment.OrderStatus)
	//  action according to
	if payment.OrderStatus == "PAID" && req.PlanPrice == float64(payment.OrderAmount) {
		customer := GetCustomerFromReq(&req)
		result := db.AddUser(&customer)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"messege": "Internal Error contact us",
				"error":   result.Error.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&customer)
	} else {
		return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
			"message": "Payment" + payment.OrderStatus,
		})
	}

}
