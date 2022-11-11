package delete

import (
	"backend/bucket"
	"backend/db"
	"backend/modal"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func DeleteImage(c *fiber.Ctx) error {
	type Image struct {
		Name string `json:"name"`
		Id   uint   `json:"id"`
		Key  string `json:"key"`
	}
	type ReqModal struct {
		FolderId uint    `json:"folder_id"`
		Image    []Image `json:"image"`
	}
	var req ReqModal
	perror := c.BodyParser(&req)
	if perror != nil {
		fmt.Println(perror.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Println("index : Started")
	for _, image := range req.Image {
		fmt.Printf("Id---- : %v", image.Id)
		_, eror := bucket.AWSS3CLIENT.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
			Bucket: aws.String("cloud-gallery-2022"),
			Key:    aws.String(image.Key),
		})
		if eror != nil {
			fmt.Printf("----2----5--------------- %v", eror.Error())
			continue
		}

		fmt.Println("Aws Deleted")

		derr2 := db.Database.DB.Unscoped().Delete(&modal.Image{}, &image.Id)
		if derr2.Error != nil {
			fmt.Printf("----2------------------- %v", derr2.Error.Error())
			continue
		}
		fmt.Println("Database Deleted")

	}

	fmt.Println("index : End")
	folder, err := db.RetriveFolder(req.FolderId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Folder Not Found",
		})
	}
	folder.Length = folder.Length - len(req.Image)
	db.Database.DB.Save(&folder)
	return c.SendStatus(200)
}
