package controllers

import (
	"blog/models"
)

type CollectController struct {
	JudgeController
}

func (c *CollectController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}
	blogList := models.QueryUserCollectWithId(models.QueryUserWithUsername(c.Loginuser.(string)))

	c.Data["Content"] = models.MakeCollectBlocks(blogList)

	c.TplName = "collect.html"
}
