package read

import (
	"backend/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetClientFolder(c *fiber.Ctx) error {
	type Req struct {
		Uid   string `json:"uid"`
		AwsId string `json:"aws_id"`
	}
	var req Req
	err := c.BodyParser(&req)
	fmt.Println(req.Uid, "_________________________")
	fmt.Println(req.AwsId, "_________________________")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid Param",
	// 		"error":   err.Error(),
	// 	})
	// }

	job, err := db.GetFolderForClient(req.AwsId, req.Uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Job Not Fount",
			"error":   err.Error(),
		})

	}
	fmt.Println(job, "_________________________")
	return c.Status(fiber.StatusOK).JSON(&job)
}
