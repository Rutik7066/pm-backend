package bucket

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var AWSS3CLIENT *s3.Client

func AwsInit() {
	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	cfg, cfgErr := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("AWS_REGION")))
	if cfgErr != nil {
		log.Printf("error: %v", cfgErr)
		return
	}

	AWSS3CLIENT = s3.NewFromConfig(cfg)
}

// func UploadImage(image []byte, ) {
// 	jpg,jpgErr :=jpeg.Decode(image)
// 	if jpgErr != nil{

// 	}
// 	uploader := manager.NewUploader(AWSS3CLIENT)
// 	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String("my-bucket"),
// 		Key:    aws.String("my-object-key"),
// 		Body:   image,
// 	})

// }
