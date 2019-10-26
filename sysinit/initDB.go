package sysinit

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.Info("数据库初始化")
	// 获取配置信息 https://beego.me/docs/mvc/controller/config.md
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")
	// set default database
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":3306)/"+mysqldb+"?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.User))

	// create table
	orm.RunSyncdb("default", false, true)
}
