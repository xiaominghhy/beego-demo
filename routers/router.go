package routers

import (
	"beego-demo/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/list", &controllers.UserController{}, "get:ListUser")

}
