package main

import (
	_ "second/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 获取配置信息 https://beego.me/docs/mvc/controller/config.md
	beego.AppConfig.String("mysqluser")
	beego.AppConfig.String("mysqlpass")
	beego.AppConfig.String("mysqlurls")
	beego.AppConfig.String("mysqldb")

	// 下载文件
	// StaticDir["/static"] = "static"
	beego.SetStaticPath("/down1", "download1")
	beego.Run()
}

/*
admin:
http://localhost:8088/qps
*/
