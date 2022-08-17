package controllers

import (
	"fmt"
	"strconv"
	"web/models"
)

type SearchController struct {
	JudgeController
}

func (c *SearchController) Get() {
	c.Data["Username"] = c.Loginuser
	keyWords := c.GetString("keyWords") //获取用户输入的字段
	pageIndex := c.GetString("page")//获取页面的page参数值,前端会发送,默认为1
	page,_ := strconv.Atoi(pageIndex)	 //转化成Int类型用于函数
	fmt.Println("pageIndex:",page)
	fmt.Println("keyWords:",keyWords)
	//page=1
	if keyWords == "" {
		empty, num := models.FindNovelWithKeyWordsOnEmpty(page)
		var novel [8]Novel
		i := 0
		for _,n := range empty{
			novel[i].Id				= n["id"].(string)
			novel[i].Name 			= n["novel_name"].(string)
			novel[i].Picture 		= n["novel_pic"].(string)
			novel[i].Writer 		= n["novel_writer"].(string)
			novel[i].Introduction 	= n["novel_introduction"].(string)
			novel[i].Genres 		= n["novel_genre"].(string)
			novel[i].State 			= n["novel_state"].(string)
			i++
		}
		c.Data["json"] 	   = map[string]interface{}{"novel":novel , "num": num, "code": 0}
		c.ServeJSON()
	}else {
		result, num := models.FindNovelWithKeyWords(keyWords, page)
		if num == 0{
			c.Data["json"] = map[string]interface{}{"code": 1, "num": 0}
			c.ServeJSON()
		}
		//返回所有匹配结果
			var novel [8]Novel
			i := 0
			for _, n := range result {
				novel[i].Id = n["id"].(string)
				novel[i].Name = n["novel_name"].(string)
				novel[i].Picture = n["novel_pic"].(string)
				novel[i].Writer = n["novel_writer"].(string)
				novel[i].Introduction = n["novel_introduction"].(string)
				novel[i].Genres = n["novel_genre"].(string)
				novel[i].State = n["novel_state"].(string)
				i++
			}
			c.Data["json"] = map[string]interface{}{"novel": novel, "num": num, "code": 0}
			c.ServeJSON()

		//c.TplName = "search.html"
	}
}
