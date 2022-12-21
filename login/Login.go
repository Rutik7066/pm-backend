package login

import (
	"backend/db"
	"backend/modal"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type logindetail struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	log.Println(c.Body())  
	log.Println(string(c.Body())) 
	var login logindetail 
	
	parErro := json.Unmarshal([]byte(c.Body()), &login)
	log.Println(login)
	if parErro != nil {
		log.Println(parErro.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"mesage": "Bad Request",
		})
	}
	var user modal.Customer
	db.Database.DB.Where("customer_email = ?", login.Email).Preload("Jobs.Images").Find(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incorrect Email",
		})
	}

	if login.Password != user.Password {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
