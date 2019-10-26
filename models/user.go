package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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
	logs.Info(r, err)
	return r
}

// 删除用户
func DeleteUser(user User) {
	o := orm.NewOrm()
	num, err := o.Delete(&user)
	logs.Info(num)
	logs.Error(err)
}

//修改用户

func UpdateUser(user User) User {
	o := orm.NewOrm()
	num, err := o.Update(&user)
	logs.Info(num)
	logs.Error(err)
	return user
}

// 传入User。返回User模型
func GetUserByName(user User) User {

	// TODO 不同的列查询，指定列查询

	o := orm.NewOrm()
	error := o.Read(&user)
	// r := o.Raw("select * from  user WHERE user_name = ?", user.UserName)

	logs.Error(error)
	return user

}
