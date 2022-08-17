package controllers

import (
	"fmt"
	"strconv"
	"web/models"
)

type Url2Controller struct {
	JudgeController
}

func (c * Url2Controller) Get() {

	pageIndex := c.GetString("page")   //获取页面的page参数值
	page, _ := strconv.Atoi(pageIndex) //转化成Int类型用于函数
	fmt.Println("pageIndex:", page)
	limitIndex := c.GetString("limit")   //获取页面的展示数量参数值
	limit, _ := strconv.Atoi(limitIndex) //转化成Int类型用于函数

	maps := models.FindUserWithPage(page,limit) //获取当前页面的数据 含limit语法
	//fmt.Println("Json:",maps)
	var user [10]models.User //因为每页只显示10个用户,所以数组大小设为10
	i := 0
	for _, n := range maps {
		user[i].Id 	  	 = n["id"].(string)
		user[i].Username = n["username"].(string)
		user[i].Genre    = n["genre"].(string)
		i++
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "count": models.QueryUserRowNum(), "data": user}
	c.ServeJSON()
}
