package controllers

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}

type LIKE struct {
	Food   string
	Watch  string
	Listen string
}

type JSONS struct {
	//必须的大写开头
	Code string
	Msg  string
	User []string `json:"user_info"` //key重命名,最外面是反引号
	Like LIKE
}

func (c *DataController) Get() {
	data := &JSONS{"100", "获取成功",
		[]string{"maple", "18"}, LIKE{"蛋糕", "电影", "音乐"}}
	response := &models.BaseResponseModel{200, "查询数据库Post", data}
	// response := models.SuccessBaseResponse(data)
	c.Data["json"] = response

	c.ServeJSON()
}
