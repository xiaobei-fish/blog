package controllers

import (
"fmt"
"strconv"
"web/models"
)

type Url1Controller struct {
	JudgeController
}

func (c * Url1Controller) Get() {

	pageIndex := c.GetString("page")   //获取页面的page参数值
	page, _ := strconv.Atoi(pageIndex) //转化成Int类型用于函数
	limitIndex := c.GetString("limit")   //获取页面的展示数量参数值
	limit, _ := strconv.Atoi(limitIndex) //转化成Int类型用于函数
	fmt.Println("pageIndex:", page)
	fmt.Println(".......",limit)

	maps := models.FindNovelWithPageLimit(page,limit) //获取当前页面的数据 含limit语法
	//fmt.Println("Json:",maps)
	var novel [10]Novel //因为每页只显示20本书 所以数组大小设为20
	i := 0
	for _, n := range maps {
		novel[i].Id = n["id"].(string)
		novel[i].Name = n["novel_name"].(string)
		novel[i].Picture = n["novel_pic"].(string)
		novel[i].Writer = n["novel_writer"].(string)
		novel[i].Introduction = n["novel_introduction"].(string)
		novel[i].Genres = n["novel_genre"].(string)
		novel[i].State = n["novel_state"].(string)
		i++
	}
	//fmt.Println("::::",map[string]interface{}{"code": 0, "msg": "", "count": models.QueryNovelRowNum(), "data": novel})
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "count": models.QueryNovelRowNum(), "data": novel}
	c.ServeJSON()
}
