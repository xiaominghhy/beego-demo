package controllers

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "file.html"
}
func (c *FileController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	_dir := "static/upload/"
	exist, err := PathExists(_dir)
	if err != nil {
		logs.Info("get dir error![%v]\n", err)
		return
	}

	if exist {
		logs.Info("has dir![%v]\n", _dir)
	} else {
		logs.Info("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			logs.Info("mkdir failed![%v]\n", err)
		} else {
			logs.Info("mkdir success!\n")
		}
	}

	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	logs.Info("上传文件名:" + h.Filename)
	c.TplName = "file.html"
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
