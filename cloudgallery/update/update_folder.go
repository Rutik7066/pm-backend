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
	if peror != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	eroor := db.Database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if eroor != nil {
		fmt.Println(eroor.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(&req)

}
