package main

import (
	"backend/bucket"
	"backend/cloudgallery/create"
	"backend/cloudgallery/credit"
	"backend/cloudgallery/delete"
	"backend/cloudgallery/read"
	"backend/cloudgallery/update"
	"backend/createaccount"
	"backend/db"
	"backend/gettoken"
	"backend/login"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	LoadEnv()
	db.ConnectDb()
	bucket.AwsInit()
}
func main() {
	fmt.Println(os.Getenv("APPID"))

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendStatus(fiber.StatusOK)
	})
	app.Post("/getcrftoken", gettoken.GetToken)
	app.Post("/getnewaccount", createaccount.NewAccount) //**
	app.Post("/demo", createaccount.Demo)                //**
	app.Post("/confirmandcreate", createaccount.CreateAccount)
	app.Post("/login", login.Login)            //**
	app.Post("/getcustomer", read.GetCustomer) //**
	app.Get("/getplan", credit.GetPlan)
	app.Post("/updatecredit", credit.UpdateCredit) //**
	app.Post("/createfolder", create.CreateFolder) //**
	app.Post("/addimage", create.AddImage)
	app.Delete("/deletefolder", delete.DeleteFolder) //**
	app.Delete("/deleteimage", delete.DeleteImage)   //**
	app.Put("/updatefolder", update.UpdateFolder)
	app.Get("/getfolder", read.GetClientFolder)
	app.Listen(":80")
}
