package sysinit

import (
	"beego-demo/models"
	"time"

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
	// set default database root:root@/orm_test?charset=utf8&loc=Asia%2FShanghai
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":3306)/"+mysqldb+"?charset=utf8", 30)
	// 鉴于 Sqlite3 的设计，存取默认都为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// register model
	orm.RegisterModel(new(models.User))

	// create table
	orm.RunSyncdb("default", false, true)
}
