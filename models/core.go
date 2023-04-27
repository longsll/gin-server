package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init()  {
	dsn := "root:123456@tcp(47.113.151.94:3306)/ch01?charset=utf8mb4&parseTime=True&loc=Local"
	DB , _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	 if err != nil {
	 	fmt.Println(err)
	}

}