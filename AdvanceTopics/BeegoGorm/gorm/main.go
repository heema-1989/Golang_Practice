package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/models"
	_ "gorm/routers"
)

var db *gorm.DB
var err error

const DSN = "heema:Simform@123@tcp(127.0.0.1:3306)/book?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Book{})
	models.Init()
	db.Create(&models.People)
	books := models.Books
	for i, _ := range books {
		db.Create(&books[i])
	}
	beego.Run()
}
