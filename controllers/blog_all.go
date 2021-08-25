package controllers

import (
	"blog/models"
	"fmt"
)

type AllBlogController struct {
	JudgeController
}

func (c *AllBlogController) Get(){
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}

	page, _ := c.GetInt("page")
	fmt.Println("page:",page)
	var blogList []models.Blog

	if page <= 0 {
		page = 1
	}
	//设置分页
	blogList, _ = models.FindAllBlogWithPage(page)
	fmt.Println(blogList)
	c.Data["PageCode"] = models.ConfigAllHomeFooterPageCode(page)
	c.Data["HasFooter"] = true

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeSearchBlocks(blogList)

	c.TplName = "blog_all.html"
}
