package routers

import (
	"beego-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/json", &controllers.DataController{})
	beego.Router("/mysql", &controllers.MysqlController{})
}
