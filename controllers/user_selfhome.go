package controllers

import (
	"fmt"
	"strconv"
	"web/models"
)

type SelfhomeController struct {
	JudgeController
}

func (c *SelfhomeController) Get() {
	userId := models.QueryUserWithUsername(c.Loginuser.(string))
	num := models.GetCollectedNumByUserId(strconv.Itoa(userId))
	number := models.GetCreatedNovelNumByWriterName(c.Loginuser.(string))

	c.Data["Username"] = c.Loginuser
	c.Data["CollectNum"] = num
	c.Data["CreateNum"] = number
	c.TplName = "selfhome.html"
}

func (c *SelfhomeController) Post() {
	username := c.Loginuser.(string)
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	fmt.Println("old:",oldPassword)
	fmt.Println("new:",newPassword)

	flag := models.TestOldPassword(oldPassword)
	if flag == 0 {
		c.Data["json"] =  map[string]interface{}{"code": 0, "message": "输入旧密码错误,请重新输入"}
		c.ServeJSON()

		return
	}else {
		models.AlterPassword(username,newPassword)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改密码成功"}
		c.ServeJSON()
	}
}