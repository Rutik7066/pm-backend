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
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
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

	// Certificate manager
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Replace with your domain
		HostPolicy: autocert.HostWhitelist("ec2-65-0-55-55.ap-south-1.compute.amazonaws.com"),
		// Folder to store the certificates
		Cache: autocert.DirCache("./certs"),
	}

	// TLS Config
	cfg := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}
	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		panic(err)
	}

	// Start server
	log.Fatal(app.Listener(ln))
}
