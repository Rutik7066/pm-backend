package login

import (
	"backend/db"
	"backend/modal"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type logindetail struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var login logindetail
	if c.BodyParser(&login) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"mesage": "Incorrect Credential",
		})
	}
	var user modal.Customer
	db.Database.DB.Where("customer_email = ?", login.Email).Preload("Jobs.Images").Find(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incorrect Email",
		})
	}

	if login.Password != user.CustomerPass {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&user)
}
