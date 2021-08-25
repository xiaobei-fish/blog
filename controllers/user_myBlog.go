package controllers

import (
	"blog/models"
	"strconv"
)

type MineController struct {
	JudgeController
}

func (c *MineController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}
	blogList := models.QueryBlogWithUserId(strconv.Itoa(models.QueryUserWithUsername(c.Loginuser.(string))))
	c.Data["Content"] = models.MakeMineBlocks(blogList)

	c.TplName = "mine.html"
}
