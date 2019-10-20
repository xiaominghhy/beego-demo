package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HelloController struct {
	beego.Controller
}
type ShortResult struct {
	UrlShort string
	UrlLong  string
}
func (this *HelloController) Get() {
	var result ShortResult
	shorturl := this.Input().Get("shorturl")

	if shorturl == "" {
		logs.Warn("shorturl为空")
		shorturl = "没有传入参数shorturl"
	} else {
		logs.Info("shorturl:+" + shorturl)
	}

	result.UrlShort = shorturl

	result.UrlLong = "haode"

	this.Data["json"] = result
	this.ServeJSON()
}
