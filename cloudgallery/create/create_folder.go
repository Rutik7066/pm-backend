package create

import (
	"backend/bucket"
	"backend/db"
	"backend/modal"
	"context"

	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

type Image struct {
	Name      string
	LocalURL  string
	ImageData multipart.File
}
type UpReq struct {
	UID    string
	Title  string
	Length int
	Image  []Image
}

func convertForm(form *multipart.Form) (req *UpReq, err error) {
	var up UpReq
	up.UID = form.Value["uid"][0]
	up.Title = form.Value["title"][0]
	baseDir := form.Value["dir"][0]
	var images []Image
	for _, fileHeader := range form.File {
		for _, file := range fileHeader {
			fmt.Printf("Type %T", file)
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			Ofile, erro := file.Open()
			if erro != nil {
				err = erro
				return
			}
			i := Image{
				Name:      file.Filename,
				LocalURL:  baseDir + "\\" + file.Filename,
				ImageData: Ofile,
			}
			images = append(images, i)
		}
	}
	up.Length = len(images)
	up.Image = images
	req = &up
	return
}

func CreateFolder(c *fiber.Ctx) error {
	form, parErr := c.MultipartForm()
	if parErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Param",
		})

	}
	upJob, parErro := convertForm(form)
	if parErro != nil {
		fmt.Println(parErro.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Printf("Uid Is %v \n", upJob.UID)
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
		uploader := manager.NewUploader(bucket.AWSS3CLIENT)
		conDi := "inline"
		conType := "image/jpeg"
		imagePath := customer.BusinessName + "/" + upJob.Title + "/" + image.Name
		// uploading to aws
		result, UploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:             aws.String("cloud-gallery-2022"),
			Key:                aws.String(imagePath),
			Body:               image.ImageData,
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
