package createaccount

import (
	"backend/db"
	"backend/modal"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewAccount(c *fiber.Ctx) error {
	var customer modal.Customer
	err := c.BodyParser(&customer)
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	customer.ValidTill = time.Now().Add(time.Hour * 8760 * 10)
	custom := db.AddUser(&customer)
	if custom.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": custom.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&customer)
}
