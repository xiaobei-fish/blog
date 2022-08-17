package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "home.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	id := models.QueryUserWithGenre(username, utils.MD5(password),"1' or genre='2")
	fmt.Println("id:", id)

	if id > 0 {
		//设置session以免密登录
		c.SetSession("loginuser", username)

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "成功    XX小说网欢迎您"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "失败,密码错误或无此用户"}
	}
	c.ServeJSON()
}
