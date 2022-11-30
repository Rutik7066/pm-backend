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

	"fmt"
	"log"
	
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("Init"))
	db.ConnectDb()
	bucket.AwsInit()
}

func main() {
	fmt.Println(os.Getenv("APPID"))
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello World  ✌")
		return c.SendString("Hello World  ✌")
	})
	app.Get("/moniter", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
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
	app.Post("/updatefolder", update.UpdateFolder)
	app.Get("/getfolder", read.GetClientFolder)
	fmt.Println("backend up & running ✌")
	app.Listen(":3000")
}
