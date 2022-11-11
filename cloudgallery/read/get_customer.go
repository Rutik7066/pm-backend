package read

import (
	"backend/db"
	"backend/modal"

	"github.com/gofiber/fiber/v2"
)

func GetCustomer(c *fiber.Ctx) error {
	type ReqModal struct {
		Uid uint
	}
	var req ReqModal
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var customer modal.Customer
	customer.ID = req.Uid
	erro := db.Database.DB.Preload("Jobs.Images").Find(&customer).Error
	if erro != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": erro.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&customer)

}
