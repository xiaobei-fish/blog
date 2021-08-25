package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type JudgeController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

//判断是否登录
func (c *JudgeController) Prepare() {
	loginuser := c.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		c.IsLogin = true
		c.Loginuser = loginuser
	} else {
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
}
