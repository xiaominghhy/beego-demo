package main

import (
	// 前面有一个下划线表示初始化该model的init方法
	_ "beego-demo/routers"
	// 系统初始化，数据库等
	_ "beego-demo/sysinit"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//设置日志信息
	logs.SetLevel(beego.LevelInformational)
	//可能会出现没有权限的情况，检查下
	logs.SetLogger("file", `{"filename":"logs/beego.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.Informational("程序开始运行")

	// 下载文件 上传文件参考 controller/file.go
	// StaticDir["/static"] = "static"
	beego.SetStaticPath("/download", "static/upload")
	beego.Run()
}

/*
admin:
http://localhost:8088/qps
*/
