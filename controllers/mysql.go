package controllers

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
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
	// 包名加结构体,什么时候需要 & 什么时候不需要& 泛型？anyType怎么定义
	data := &models.BaseResponseModel{"200", "查询数据库", models.User{"eric", 24}}
	c.Data["json"] = data
	c.ServeJSON()
}
