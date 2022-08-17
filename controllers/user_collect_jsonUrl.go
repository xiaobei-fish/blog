package controllers

import (
	"fmt"
	"strconv"
	"web/models"
)

type CollectController struct {
	JudgeController
}

type Novel struct {
	Id         		string
	Name     		string
	Genres       	string
	Introduction    string
	Writer     		string
	Picture			string
	State			string
}

func (c *CollectController) Get() {

	pageIndex := c.GetString("page")//获取页面的page参数值
	page,_ := strconv.Atoi(pageIndex)	 //转化成Int类型用于函数
	fmt.Println("pageIndex:",page)
	//得到当前用户在数据库中的id
	userid := models.QueryUserWithUsername(c.Loginuser.(string))
	//得到当前用户的收藏夹数据
	maps,num := models.FindNovelWithCollectId(page, strconv.Itoa(userid))
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
	//c.Data["json"]     = novel
	//c.TplName = "collect.html"
	c.ServeJSON()
}