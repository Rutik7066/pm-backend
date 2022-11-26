package createaccount

import (
	"backend/db"
	"backend/modal"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Demo(c *fiber.Ctx) error {
	var customer modal.Customer
	perror := c.BodyParser(&customer)
	if perror != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	customer.Credit = 50
	customer.ValidTill = time.Now().Add(time.Hour * 24 * 3)
	dberror := db.Database.DB.Create(&customer).Error
	if dberror != nil {
		fmt.Println(dberror)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(&customer)

}
