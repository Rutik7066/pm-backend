package read

import (
	"backend/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetClientFolder(c *fiber.Ctx) error {
	uid := c.Query("uid")
	awsid := c.Query("aws_id")
	log.Println(uid, awsid)
	job, err := db.GetFolderForClient(awsid, uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Job Not Fount",
			"error":   err.Error(),
		})

	}
	fmt.Println(job, "_________________________")
	return c.Status(fiber.StatusOK).JSON(&job)
}
