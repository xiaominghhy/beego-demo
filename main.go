package main

import (
	// 前面有一个下划线表示初始化init方法
	_ "beego-demo/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 获取配置信息 https://beego.me/docs/mvc/controller/config.md
	beego.AppConfig.String("mysqluser")
	beego.AppConfig.String("mysqlpass")
	beego.AppConfig.String("mysqlurls")
	beego.AppConfig.String("mysqldb")

	// 下载文件 上传文件参考 controller/file.go
	// StaticDir["/static"] = "static"
	beego.SetStaticPath("/down1", "static/upload")
	beego.Run()
}

/*
admin:
http://localhost:8088/qps
*/
