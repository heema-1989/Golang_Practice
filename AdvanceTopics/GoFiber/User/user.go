package User

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "heema:Simform@123@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local"

type Users struct {
	gorm.Model
	Firstname string `json:"firstname`
	Lastname  string `json:"lastname"`
	Email     string `json:email`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to database")
	}
	DB.AutoMigrate(&Users{})
}
func GetUsers(c *fiber.Ctx) error {
	var users []Users
	DB.Find(&users)
	return c.JSON(&users)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user Users
	DB.Find(&user, id)
	return c.JSON(&user)
}
func SaveUser(c *fiber.Ctx) error {
	user := new(Users)
	if err = c.BodyParser(user); err != nil {
		return c.Status(5000).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)
}
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(Users)  //this will create a new instance of struct Users
	DB.First(&user, id) //here we can also do DB.First(&Users{},id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err = c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(Users)
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err = c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Delete(&user)
	return c.JSON(&user)
}
