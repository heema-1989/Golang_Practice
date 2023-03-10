package routers

import (
	"github.com/astaxie/beego"
	"gorm/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/createP", &controllers.TestController{})
}
