package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sitepointgoapp/models"
	_ "sitepointgoapp/routers"
)

func main() {
	orm.Debug = true
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	fmt.Println(err)

	err = orm.RegisterDataBase("default", "mysql", "heema:Simform@123@/beego?")
	orm.RegisterModel(new(models.Users))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("Failed to create tables")
	} else {
		fmt.Println("Success")
	}

	fmt.Println(err)
	beego.Run()
}
