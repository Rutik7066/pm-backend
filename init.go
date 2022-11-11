package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const CRFURL = "https://sandbox.cashfree.com/pg"

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APPID"))
}
