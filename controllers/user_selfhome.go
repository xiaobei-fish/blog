package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type SelfHomeController struct {
	JudgeController
}

func (c *SelfHomeController) Get(){
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}
	//导入收藏信息
	username := c.Loginuser.(string)
	info := models.QueryUserInfoWithUsername(username)
	fmt.Println(info)

	for _, i := range info{
		c.Data["userId"] = i.Id
		c.Data["Username"] = i.Username
		c.Data["Phone"] = i.Phone
	}

	c.Data["Img"] = models.QueryUserHeadImgWithUsername(username)

	c.TplName = "selfhome.html"
}

func (c *SelfHomeController) Post() {
	c.EnableRender = false
	username := c.GetString("username")

	fmt.Println("new username:",username)
	_, err := models.UpdateInfo(strconv.Itoa(models.QueryUserWithUsername(c.Loginuser.(string))),username)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败"}
	}else{
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功"}
		c.JudgeController.DelSession("loginuser")
		c.SetSession("loginuser", username)
		c.JudgeController.Loginuser = c.GetSession("loginuser")
	}
	c.ServeJSON()
}
