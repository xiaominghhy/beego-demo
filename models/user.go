package models

import (
	"fmt"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	//必须的大写开头
	Id          int
	UserName    string `orm:"size(20);null"`
	Password    string `orm:"size(32);null"`
	Tel         string `orm:"size(11);null"`
	Address     string `orm:"null"`
	PaperNumber int    `orm:"null"`
}

//下面可以写一些关于User得增删该查操作，所有的方法都是基于Model这个概念得。
/*
有啥就使用啥，查询的时候，哪个字段存在就查询哪个字段
*/

//添加用户
func AddUser(user User) int64 {
	logs.Info("新建用户", user)
	o := orm.NewOrm()
	r, err := o.Insert(&user)
	logs.Info(r)
	if err != nil {
		logs.Error(err)
	}
	return r
}

// 删除用户
func DeleteUser(user User) {
	o := orm.NewOrm()
	num, err := o.Delete(&user)
	logs.Info(num)
	if err != nil {
		logs.Error(err)
	}
}

//修改用户

func UpdateUser(user User) User {
	o := orm.NewOrm()
	num, err := o.Update(&user)
	logs.Info(num)
	if err != nil {
		logs.Error(err)
	}
	return user
}

// 高级查询https://beego.me/docs/mvc/model/query.md#one

// 传入User。返回User模型
func GetUserByName(user User) User {
	//  不同的列查询，指定列查询
	o := orm.NewOrm()
	error := o.Read(&user, "UserName")
	if error != nil {
		logs.Error(error)
	}
	return user
}

// 传入User。返回User模型
func GetUserById(user User) User {

	o := orm.NewOrm()
	// 主键查询
	error := o.Read(&user)
	// r := o.Raw("select * from  user WHERE user_name = ?", user.UserName)
	if error != nil {
		logs.Error(error)
	}
	return user

}

func ListUser(limit, page int) (int, int) {

	fmt.Println(limit)
	fmt.Println(page)
	return limit, page
	// orm := orm.NewOrm()
	// // 获取表句柄
	// qs := orm.QueryTable("user")
	// //select * from student limit 5 ;
	// n,err := qs.Limit(5) .All(&user)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// logs.Error(n)
	// return user
	// return limit,page
}
