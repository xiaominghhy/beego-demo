package main

import (
	// 前面有一个下划线表示初始化该model的init方法
	_ "beego-demo/routers"
	// 系统初始化，数据库等
	_ "beego-demo/sysinit"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//设置日志信息
	logs.SetLevel(logs.LevelInformational)
	//可能会出现没有权限的情况，检查下
	logs.SetLogger("file", `{"filename":"logs/beego.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.Informational("程序开始运行")
	// 日志打印的另一种方式，但是这种方式已经被官方放弃了，建议使用上面那种
	logs.Info("日志的另一种方式,相比较上面好像文件与行数更加准确")

	// 下载文件 上传文件参考 controller/file.go
	// StaticDir["/static"] = "static"
	beego.SetStaticPath("/download", "static/upload")
	beego.Run()
}

/*
admin:
http://localhost:8088/qps
*/
