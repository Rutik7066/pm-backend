package modal

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID               uint `gorm:"primary Key;autoIncrement" json:"uid"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	CustomerName     string         `json:"customer_name"`
	CustomerPhone    string         `json:"customer_phone"`
	CustomerAltPhone string         `json:"customer_alt_phone"`
	CustomerEmail    string         `json:"customer_email"`
	CustomerPass     string         `json:"password"`
	BusinessName     string         `json:"business_name"`
	BusinessAddress  string         `json:"business_address"`
	FbId             string         `json:"fb_id"`
	SnapId           string         `json:"snap_id"`
	InstaId          string         `json:"insta_id"`
	Web              string         `json:"web"`
	Youtube          string         `json:"youtube"`
	// youtube channel
	IpAddress string    `json:"ip_address"`
	PlanPrice float64   `json:"plan_price"`
	Credit    int       `json:"credit"`
	ValidTill time.Time `json:"valid_till"`
	Jobs      []*Job    `json:"jobs"`
}
type Job struct {
	ID         uint `gorm:"primary Key;autoIncrement" json:"id"`
	CustomerID uint
	AwsId      string   `json:"aws_id"`
	Status     int      `json:"status"`
	Length     int      `json:"length"`
	Images     []*Image `json:"images"`
}
type Image struct {
	ID         uint `gorm:"primary Key;autoIncrement" json:"id"`
	JobID      uint
	Key        *string `json:"key"`
	Name       string  `json:"name"`
	Localurl   string  `json:"local_url"`
	BucketUrl  string  `json:"bucket_url"`
	IsSelected bool    `json:"is_selected"`
}
