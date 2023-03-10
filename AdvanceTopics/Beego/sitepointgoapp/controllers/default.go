package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sitepointgoapp/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (main *MainController) AddUser() {
	main.TplName = "default/login.html"
}
func (main *MainController) AddedUser() {

	main.TplName = "default/welcome.html"
	main.Data["Name"] = main.GetString("name")
	main.Data["Email"] = main.GetString("email")
	o := orm.NewOrm()
	e := o.Using("default")
	if e != nil {
		return
	}
	/*qs := o.QueryTable(models.Users{})
	del, _ := qs.Filter("id", 6).Update(orm.Params{"id": 1})
	fmt.Println("User id updated ", del)*/
	user := models.Users{}
	err := main.ParseForm(&user)
	fmt.Println("ZERO")
	if err != nil {
		fmt.Println("Couldn't parse the form: Reason ", err)
	} else {
		main.Data["Users"] = user
		if main.Ctx.Input.Method() == "POST" {
			fmt.Println("FIRST")
			user.Username = main.GetString("username")
			user.Name = main.GetString("name")
			user.Email = main.GetString("email")
			user.Password = main.GetString("password")
			fmt.Println(user)
		}
	}

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}

func (main *MainController) DeleteUser() {
	main.TplName = "default/delete.html"
	userId, _ := strconv.Atoi(main.Ctx.Input.Param(":id"))
	main.Data["id"] = userId
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.Users{})
	del, err := qs.Filter("id", userId).Delete()
	if err != nil {
		fmt.Println("Failed to delete user with id: ", userId, " Reason: ", err)
	}
	fmt.Println("Successfully deleted user with id: ", del)
}
func (main *MainController) UpdateUser() {
	main.TplName = "default/update.html"
	userId, err := strconv.Atoi(main.Ctx.Input.Param(":id"))
	fmt.Println(userId)
	if err != nil {
		fmt.Println("Enter proper id")
	}
	o := orm.NewOrm()
	o.Using("default")

	main.Data["Id"] = userId
	user := models.Users{Id: userId}
	var flag bool
	if o.Read(&user) == nil {
		flag = true
		main.Data["Name"] = user.Name
		main.Data["Username"] = user.Username
		main.Data["Email"] = user.Email
		main.Data["Password"] = user.Password
	} else if o.Read(&user) == orm.ErrNoRows {
		flag = false
		fmt.Println("Record not found")
	}
	main.Data["flag"] = flag
	fmt.Println(user)
	fmt.Println(flag)
	fmt.Println(user.Name)
}

func (main *MainController) UpdatedUser() {
	main.TplName = "default/updated.html"

	o := orm.NewOrm()
	o.Using("default")
	userId, _ := strconv.Atoi(main.GetString("id"))
	user := models.Users{Id: userId}
	main.ParseForm(&user)
	fmt.Println(user.Id, userId)
	flash := beego.NewFlash()
	fmt.Println(user)
	fmt.Println("first")
	if main.Ctx.Input.Method() == "POST" {
		fmt.Println("second")
		err := o.Read(&user)
		if err == nil {
			fmt.Println("THIRD")

			user.Username = main.GetString("username")
			user.Name = main.GetString("name")
			user.Email = main.GetString("email")
			user.Password = main.GetString("password")
			fmt.Println(user)
			if num, err := o.Update(&user); err == nil {
				flash.Notice("Record was updated")
				flash.Store(&main.Controller)
				beego.Info("Record Was Updated. ", num)

			}
		} else if err == orm.ErrNoRows {
			fmt.Println("Record not found")
		}
	}
	main.Data["Name"] = user.Name
	main.Data["Username"] = user.Username
	main.Data["Email"] = user.Email
	main.Data["Password"] = user.Password
}
