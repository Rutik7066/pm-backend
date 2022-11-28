package update

import (
	"backend/db"
	"backend/modal"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateFolder(c *fiber.Ctx) error {
	var req modal.Job
	peror := c.BodyParser(&req)
	temp := c.Body()
	fmt.Println(string(temp))
	if peror != nil {
		return c.Status(fiber.StatusBadRequest).JSON(peror.Error())
	}
	eroor := db.Database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if eroor != nil {
		fmt.Println(eroor.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(&req)

}
