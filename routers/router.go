package routers

import (
	"beego-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/list", &controllers.UserController{},"get:ListUser")

}
