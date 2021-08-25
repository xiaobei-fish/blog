package controllers

import "blog/models"

type HomeController struct {
	JudgeController
}

func (c *HomeController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}
	//热搜榜
	blogList := models.QueryTop5Blog()
	c.Data["Content"] = models.MakeHomeBlocks(blogList)

	c.TplName = "home.html"
}
