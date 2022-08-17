package controllers

import (
	"fmt"
	"log"
	"strconv"
	"web/models"
)

type UrlController struct {
	JudgeController
}

func (c * UrlController) Get() {

	pageIndex := c.GetString("page")   //获取页面的page参数值
	page, _ := strconv.Atoi(pageIndex) //转化成Int类型用于函数
	fmt.Println("pageIndex:", page)

	maps,num := models.FindNovelWithPage(page,c.Loginuser.(string)) //获取当前页面的数据 含limit语法
	//fmt.Println("Json:",maps)
	var novel [8]Novel//因为每页只显示8本书 所以数组大小设为8
	i := 0
	for _,n := range maps{
		novel[i].Id				= n["id"].(string)
		novel[i].Name 			= n["novel_name"].(string)
		novel[i].Picture 		= n["novel_pic"].(string)
		novel[i].Writer 		= n["novel_writer"].(string)
		novel[i].Introduction 	= n["novel_introduction"].(string)
		novel[i].Genres 		= n["novel_genre"].(string)
		novel[i].State 			= n["novel_state"].(string)
		i++
	}
	c.Data["json"] 	   = map[string]interface{}{"novel":novel , "num": num}
	c.ServeJSON()
}

func (c *UrlController) Post()  {

	Id := c.GetString("id")
	fmt.Println("删除的创作小说id:", Id)

	_, err := models.DeleteCreatedNovel(Id)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/writer", 302)
}

