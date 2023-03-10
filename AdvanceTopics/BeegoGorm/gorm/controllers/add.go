package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/models"
)

type TestController struct {
	beego.Controller
}

var db *gorm.DB
var err error

const DSN = "heema:Simform@123@tcp(127.0.0.1:3306)/book?charset=utf8&parseTime=True&loc=Local"

func (main *TestController) Post() {
	var req models.Person
	json.Unmarshal(main.Ctx.Input.RequestBody, &req)
	fmt.Println("Hello", main.Ctx.Input.RequestBody)
	fmt.Println(req)
	main.TplName = "post.tpl"
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")
	p := db.Create(&req)
	if p.Error != nil {
		main.Data["json"] = "Cannot create person"
		return
	}
	main.Data["json"] = req
	main.ServeJSON()
}
