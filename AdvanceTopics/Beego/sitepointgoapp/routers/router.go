package routers

import (
	"github.com/astaxie/beego"
	"sitepointgoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/welcome", &controllers.MainController{}, "get:AddUser")
	beego.Router("/register", &controllers.MainController{}, "post:AddedUser")
	beego.Router("/delete/:id", &controllers.MainController{}, "*:DeleteUser")
	beego.Router("/update/:id", &controllers.MainController{}, "get:UpdateUser")
	beego.Router("/updated", &controllers.MainController{}, "post:UpdatedUser")
	//beego.Router("/login", &controllers.MainController{}, "get:LoginUser")
}
