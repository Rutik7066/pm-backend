package create

import (
	"backend/bucket"
	"backend/db"
	"log"

	"backend/modal"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func CreateFolder(c *fiber.Ctx) error {
	type Image struct {
		Name     string `json:"name"`
		LocalURL string `json:"localURL"`
		Base64   string `json:"base64"`
	}
	type UpReq struct {
		UID    string  `json:"uid"`
		Title  string  `json:"title"`
		Length int     `json:"length"`
		Image  []Image `json:"job"`
	}

	var upJob UpReq
	if parErr := c.BodyParser(&upJob); parErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Param",
		})

	}
	fmt.Println(upJob.UID)
	customer := db.GetCustomer(upJob.UID)
	if customer.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorizresd",
		})
	}
	if customer.Credit < upJob.Length {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Insufficient Credit",
		})
	}

	var failed []*modal.Image
	var success []*modal.Image

	for _, image := range upJob.Image {
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
		imagePath := customer.BusinessName + "/" + upJob.Title + "/" + image.Name
		// uploading to aws
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
			log.Printf("Image Upload Failed %v", image.Name)
			continue
		}
		// listing succes ent
		i := modal.Image{
			Name:       image.Name,
			Localurl:   image.LocalURL,
			Key:        result.Key,
			BucketUrl:  result.Location,
			IsSelected: false,
		}
		success = append(success, &i)
		log.Printf("Image Uploaded %v", image.Name)

	}
	tempJob := modal.Job{
		AwsId:  upJob.Title,
		Status: 0,
		Length: len(success),
		Images: success,
	}
	customer.Credit = customer.Credit - len(success)
	customer.Jobs = append(customer.Jobs, &tempJob)
	result := db.UpdateUser(&customer)
	if result.Error != nil {
		result := db.UpdateUser(&customer)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update Record",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"failed":        &failed,
		"failed_length": len(failed),
		"credit":        &customer.Credit,
		"folder":        &tempJob,
	})
}
