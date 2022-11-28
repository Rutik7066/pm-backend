package delete

import (
	"backend/bucket"
	"backend/db"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func DeleteFolder(c *fiber.Ctx) error {
	type ReqModal struct {
		Id uint `json:"folder_id"`
	}
	var req ReqModal
	if perr := c.BodyParser(&req); perr != nil {
		fmt.Println(perr)

		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Println(1)
	folder, err := db.RetriveFolder(req.Id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Folder Not Found",
		})
	}
	fmt.Println(2)
	for _, image := range folder.Images {
		log.Println(&image.Key, "=====================================")
		_, eror := bucket.AWSS3CLIENT.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
			Bucket: aws.String("cloud-gallery-2022"),
			Key:    &image.Key,
		})
		if eror != nil {
			continue
		}
	}
	fmt.Println(3)
	eriir := db.DeleteFolder(req.Id)
	if eriir != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": eriir.Error(),
		})
	}
	fmt.Println(4)
	return c.SendStatus(fiber.StatusOK)
}
