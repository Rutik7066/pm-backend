package create

import (
	"backend/bucket"
	"backend/db"
	"backend/modal"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AddImage(c *fiber.Ctx) error {
	type Image struct {
		Name     string `json:"name"`
		LocalURL string `json:"localURL"`
		Base64   string `json:"base64"`
	}
	type UpReq struct {
		UID    string  `json:"uid"`
		AwsID  string  `json:"aws_id"`
		Length int     `json:"length"`
		Image  []Image `json:"image"`
	}
	var req UpReq
	fmt.Println("____________________------------------", req.AwsID)
	err := c.BodyParser(&req)
	if err != nil {
		fmt.Println(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}
	customer := db.GetCustomer(req.UID)

	if customer.ID == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if customer.Credit < req.Length {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Insufficient Credit",
		})
	}
	var failed []*modal.Image
	var success []*modal.Image
	for _, image := range req.Image {
		// decoding
		data, err := base64.StdEncoding.DecodeString(image.Base64)
		if err != nil {
			i := modal.Image{
				Name:       image.Name,
				Localurl:   image.LocalURL,
				IsSelected: false,
			}
			failed = append(failed, &i)
			continue
		}
		// converting to image
		imageData := bytes.NewBuffer(data)
		uploader := manager.NewUploader(bucket.AWSS3CLIENT)
		conDi := "inline"
		conType := "image/jpeg"
		imagePath := customer.BusinessName + "/" + req.AwsID + "/" + image.Name
		result, UploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:             aws.String("cloud-gallery-2022"),
			Key:                aws.String(imagePath),
			Body:               imageData,
			ContentDisposition: &conDi,
			ContentType:        &conType,
		})
		fmt.Println(result.Location)
		if UploadErr != nil {
			i := modal.Image{
				Name:       image.Name,
				Localurl:   image.LocalURL,
				IsSelected: false,
			}
			failed = append(failed, &i)
			continue
		}
		// listing succes ent
		i := modal.Image{
			Name:       image.Name,
			Localurl:   image.LocalURL,
			BucketUrl:  result.Location,
			IsSelected: false,
			Key:        &req.AwsID,
		}
		success = append(success, &i)

	}

	var folder modal.Job
	erroD := db.Database.DB.Where("customer_id = ? AND aws_id = ?", req.UID, req.AwsID).Preload(clause.Associations).First(&folder)
	if erroD.Error != nil {
		fmt.Println(erroD.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Record 1",
		})
	}
	fmt.Println(folder.AwsId)
	fmt.Println(folder.CustomerID)
	fmt.Println(folder.ID)
	fmt.Println(folder.Images)
	fmt.Println(folder.Status)
	folder.Images = append(folder.Images, success...)
	folder.Length = len(folder.Images)
	db.Database.DB.Save(&folder)
	customer.Credit = customer.Credit - len(success)
	result := db.UpdateUser(&customer)
	if result.Error != nil {
		result := db.UpdateUser(&customer)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update Record 2",
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"failed":        &failed,
		"failed_length": len(failed),
		"credit":        &customer.Credit,
		"job":           &success,
	})
}
