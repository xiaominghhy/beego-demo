package models

// import (
// 	"fmt"

// 	"github.com/astaxie/beego/logs"
// 	"github.com/astaxie/beego/orm"
// 	_ "github.com/go-sql-driver/mysql" // import your used driver
// )

// // Model Struct
// type User struct {
// 	Id          int
// 	UserName    string `orm:"size(20);null"`
// 	Password    string `orm:"size(32);null"`
// 	Tel         string `orm:"size(11);null"`
// 	Address     string `orm:"null"`
// 	PaperNumber int    `orm:"null"`
// 	// create_time  orm.DateTimeField
// }

// func init() {
// 	// set default database
// 	orm.RegisterDataBase("default", "mysql", "root:authentication_string@tcp(129.28.19.203:3306)/beego?charset=utf8", 30)

// 	// register model
// 	orm.RegisterModel(new(User))

// 	// create table
// 	orm.RunSyncdb("default", false, true)
// }

// func main() {
// 	orm.Debug = true
// 	o := orm.NewOrm()
// 	logs.Info("开始插入数据")
// 	user := User{UserName: "eric"}

// 	// insert
// 	id, err := o.Insert(&user)
// 	logs.Info(user)
// 	fmt.Printf("ID: %d, ERR: %v\n", id, err)

// 	// update
// 	user.UserName = "yangzhao"
// 	num, err := o.Update(&user)
// 	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

// 	// read one
// 	u := User{Id: user.Id}
// 	err = o.Read(&u)
// 	logs.Info(user)
// 	fmt.Printf("ERR: %v\n", err)

// 	user2 := User{Id: 8}
// 	err = o.Read(&user2)
// 	logs.Info(user2)
// 	user2.Password = "sssssss"
// 	num, err = o.Update(&user2)
// 	err = o.Read(&user2)
// 	logs.Info(user2)

// 	// delete
// 	// num, err = o.Delete(&u)
// 	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
// }
