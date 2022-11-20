package db

import (
	"backend/modal"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBINSTANCE struct {
	DB *gorm.DB
} 


var Database DBINSTANCE

func ConnectDb() {

	dsn := "host=database-2.c9aejhicjxyn.ap-south-1.rds.amazonaws.com user=postgres password=Rutik123 dbname=mydb port=5432 sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database____________________________________. \n")
	}

	if err := db.AutoMigrate(&modal.Customer{}, &modal.Job{}, &modal.Image{}); err != nil {
		fmt.Println("____________________________________________________________________\n", err)
	}

	Database = DBINSTANCE{
		DB: db,
	}
}

func AddUser(user *modal.Customer) *gorm.DB {
	return Database.DB.Create(&user)
}

// /
func AddCredit(uid string, credit int) *modal.Customer {
	var user modal.Customer
	Database.DB.First(&user, uid)
	Database.DB.Model(&user).Update("credit", (user.Credit + credit))
	return &user
}

func UpdateUser(user *modal.Customer) *gorm.DB {
	return Database.DB.Save(&user)
}

func GetCustomer(uid string) modal.Customer {
	var customer modal.Customer
	Database.DB.Where("id = ?", uid).Preload("images").First(&customer)
	return customer
}

func GetFolderForClient(awsid string, uid string) (job modal.Job, err error) {
	var tjob modal.Job
	result := Database.DB.Where("customer_id = ? AND aws_id = ?", uid, awsid).Preload("Images").First(&tjob)
	if result.Error != nil {
		err = result.Error
		return
	}
	job = tjob
	return
}

func GetFullCustomer(uid uint) (cust *modal.Customer, ero error) {
	var customer modal.Customer
	customer.ID = uid
	err := Database.DB.Preload("Jobs.Images").Find(&customer).Error
	if err != nil {
		ero = err
		return
	}
	cust = &customer
	return
}

func GetCustomerAllFolder(uid string) (job interface{}, err error) {
	var folder interface{}
	result := Database.DB.Where("id = ?", uid).Preload("jobs").Preload("images").Find(&folder)
	if result.Error != nil {
		err = result.Error
		return
	}
	job = folder
	return
}

func DeleteFolder(id uint) (eror error) {
	var folder modal.Job
	folder.ID = id
	err := Database.DB.Select(clause.Associations).Delete(&folder).Error
	if err != nil {
		eror = err
		return
	}
	return
}

func RetriveFolder(id uint) (folderData *modal.Job, erro error) {
	var folder modal.Job
	folder.ID = id
	err := Database.DB.Preload("Images").First(&folder).Error
	if err != nil {
		erro = err
		return
	}
	log.Println(len(folder.Images), "-----------ooooooooooooooooo--------------")
	folderData = &folder
	return
}
