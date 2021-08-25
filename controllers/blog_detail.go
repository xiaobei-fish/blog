package controllers

import (
	"blog/models"
	"fmt"
)

type DetailController struct {
	JudgeController
}

func (c *DetailController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}

	blogId := c.GetString("id")
	fmt.Println("blog_id:" + blogId)

	//点击就要加1浏览量
	_, _ = models.UpdateRead(blogId)

	blog := models.QueryBlogWithBlogId(blogId)

	for _, b := range blog {
		c.Data["Id"] = b.Id
		c.Data["Title"] = b.Title
		c.Data["Author"] = b.Author
		c.Data["Tag"] = b.Tag
		c.Data["Like"] = b.Like
		c.Data["Read"] = b.Read
		c.Data["Content"] = b.Content
	}

	c.TplName = "detail.html"
}


