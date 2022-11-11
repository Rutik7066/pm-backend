package credit

import (
	"backend/db"
	"backend/helper"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateCredit(c *fiber.Ctx) error {
	type creditReq struct {
		OrderId  string `json:"order_id"`
		Uid      string `json:"uid"`
		PlanName string `json:"plan_name"`
	}
	var req creditReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messege": "Invalid Param",
		})
	}
	paymentDetail, err := helper.GetPaymentDetail(req.OrderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messege": "Internal Error",
			"error":   err.Error(),
		})
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
	amt, _ := strconv.Atoi(req.PlanName)
	if paymentDetail.OrderStatus == "PAID" && paymentDetail.OrderAmount == amt {
		user := db.AddCredit(req.Uid, credit)
		fmt.Println(amt)
		fmt.Println(user)
		return c.Status(fiber.StatusOK).JSON(&user)
	} else {
		return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
			"messege": "Payment" + paymentDetail.OrderStatus,
		})
	}

}
