package controllers

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	userName := this.GetString("userName")
	user := models.User{UserName: userName}
	logs.Info(user)
	userR := models.AddUser(user)
	logs.Info(userR)

	// 返回json数据
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *UserController) Delete() {
	id, err := this.GetInt("id")
	if err != nil {
		logs.Error(err)
	}

	user := models.User{Id: id}
	logs.Info(user)
	models.DeleteUser(user)
	// logs.Info(userR)

	// 返回json数据
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *UserController) Put() {
	id, err := this.GetInt("id")
	if err != nil {
		logs.Error(err)
	}
	userName := this.GetString("userName")
	user := models.User{Id: id, UserName: userName}

	logs.Info(user)
	userR := models.UpdateUser(user)
	logs.Info(userR)

	// 返回json数据
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *UserController) Get() {
	id, err := this.GetInt("id")
	if err != nil {
		logs.Error(err)
	}

	userName := this.GetString("username")
	user := models.User{Id: id, UserName: userName}
	user = models.GetUserByName(user)
	logs.Info(user)
	// 返回json数据
	this.Data["json"] = user
	this.ServeJSON()
}
