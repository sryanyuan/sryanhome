package routers

import (
	"github.com/astaxie/beego"
	"sryanhome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
}
