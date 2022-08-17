package controllers

import "web/models"

type UserAlterPasswordController struct {
	JudgeController
}

func (c *UserAlterPasswordController) Post() {
	username := c.Loginuser.(string)
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")

	//查数据库看看旧密码是否正确
	flag := models.TestOldPassword(oldPassword)
	if flag == 0 {
		c.Data["json"] =  map[string]interface{}{"code": 0, "message": "输入旧密码错误,请重新输入"}
		c.ServeJSON()

		return
	}else {
		models.AlterPassword(username,newPassword)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改密码成功"}
		//删除用户的session
		c.DelSession("loginuser")

		c.ServeJSON()
	}
}
