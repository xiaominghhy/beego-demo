package controllers

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MysqlController struct {
	beego.Controller
}

type UserExtend struct {
	Food   string
	Watch  string
	Listen string
}

func (c *MysqlController) Get() {
	// 获取参数
	id, err1 := c.GetInt("id")
	logs.Info(id, err1)
	// 包名加结构体,什么时候需要 & 什么时候不需要& 泛型？anyType怎么定义
	orm.Debug = true
	o := orm.NewOrm()
	//  read one
	user := models.User{Id: id}
	err := o.Read(&user)
	logs.Info(user, err)

	data := &models.BaseResponseModel{"200", "查询数据库", user}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *MysqlController) Post() {
	// 包名加结构体,什么时候需要 & 什么时候不需要& 泛型？anyType怎么定义
	data := &models.BaseResponseModel{"200", "查询数据库Post", models.User{Id: 6}}
	c.Data["json"] = data
	// models.addUser()
	c.ServeJSON()
}
