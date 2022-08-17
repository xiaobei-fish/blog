package controllers

import (
	//"fmt"
	"github.com/astaxie/goredis"
	"web/models"
)

type HomeController struct {
	JudgeController
}

func (c *HomeController) Get() {
	var client goredis.Client
	client.Addr="127.0.0.1:6379"

	P1,err1 := client.Get("P1")
	if err1 != nil {
		panic(err1)
	}

	P2,err2 := client.Get("P2")
	if err2 != nil {
		panic(err2)
	}

	P3,err3 := client.Get("P3")
	if err3 != nil {
		panic(err3)
	}

	P4,err4 := client.Get("P4")
	if err4 != nil {
		panic(err4)
	}

	P5,err5 := client.Get("P5")
	if err5 != nil {
		panic(err5)
	}

	P6,err6 := client.Get("P6")
	if err6 != nil {
		panic(err6)
	}
	i1 := models.ShowPictureOnHomeByAdmin(string(P1))
	i2 := models.ShowPictureOnHomeByAdmin(string(P2))
	i3 := models.ShowPictureOnHomeByAdmin(string(P3))
	i4 := models.ShowPictureOnHomeByAdmin(string(P4))
	i5 := models.ShowPictureOnHomeByAdmin(string(P5))
	//转化为string，并且传递到html中
	//轮番图
	c.Data["P1"]  = i1[0]["novel_pic"].(string)
	c.Data["P2"]  = i2[0]["novel_pic"].(string)
	c.Data["P3"]  = i3[0]["novel_pic"].(string)
	c.Data["P4"]  = i4[0]["novel_pic"].(string)
	c.Data["P5"]  = i5[0]["novel_pic"].(string)
	c.Data["P6"]  = string(P6)
	c.Data["ID1"] = string(P1)
	c.Data["ID2"] = string(P2)
	c.Data["ID3"] = string(P3)
	c.Data["ID4"] = string(P4)
	c.Data["ID5"] = string(P5)

	//排行榜第一
	No1 := models.GetNo1Novel()
	c.Data["id"]  		   = No1[0]["id"].(string)
	c.Data["pic"] 		   = No1[0]["novel_pic"].(string)
	c.Data["name"]		   = No1[0]["novel_name"].(string)
	c.Data["genre"] 	   = No1[0]["novel_genre"].(string)
	c.Data["writer"] 	   = No1[0]["novel_writer"].(string)
	c.Data["introduction"] = No1[0]["novel_introduction"].(string)
	//排行榜第二至第十四
	Rank := models.GetRankNovels()
	c.Data["id2"] 		   = Rank[0]["id"].(string)
	c.Data["name2"]  	   = Rank[0]["novel_name"].(string)
	c.Data["genre2"] 	   = Rank[0]["novel_genre"].(string)
	c.Data["writer2"] 	   = Rank[0]["novel_writer"].(string)

	c.Data["id3"] 		   = Rank[1]["id"].(string)
	c.Data["name3"]  	   = Rank[1]["novel_name"].(string)
	c.Data["genre3"] 	   = Rank[1]["novel_genre"].(string)
	c.Data["writer3"] 	   = Rank[1]["novel_writer"].(string)

	c.Data["id4"] 		   = Rank[2]["id"].(string)
	c.Data["name4"]  	   = Rank[2]["novel_name"].(string)
	c.Data["genre4"] 	   = Rank[2]["novel_genre"].(string)
	c.Data["writer4"] 	   = Rank[2]["novel_writer"].(string)

	c.Data["id5"] 		   = Rank[3]["id"].(string)
	c.Data["name5"]  	   = Rank[3]["novel_name"].(string)
	c.Data["genre5"] 	   = Rank[3]["novel_genre"].(string)
	c.Data["writer5"] 	   = Rank[3]["novel_writer"].(string)

	c.Data["id6"] 		   = Rank[4]["id"].(string)
	c.Data["name6"]  	   = Rank[4]["novel_name"].(string)
	c.Data["genre6"] 	   = Rank[4]["novel_genre"].(string)
	c.Data["writer6"] 	   = Rank[4]["novel_writer"].(string)

	c.Data["id7"] 		   = Rank[5]["id"].(string)
	c.Data["name7"]  	   = Rank[5]["novel_name"].(string)
	c.Data["genre7"] 	   = Rank[5]["novel_genre"].(string)
	c.Data["writer7"] 	   = Rank[5]["novel_writer"].(string)

	c.Data["id8"] 		   = Rank[6]["id"].(string)
	c.Data["name8"]  	   = Rank[6]["novel_name"].(string)
	c.Data["genre8"] 	   = Rank[6]["novel_genre"].(string)
	c.Data["writer8"] 	   = Rank[6]["novel_writer"].(string)

	c.Data["id9"] 		   = Rank[7]["id"].(string)
	c.Data["name9"]  	   = Rank[7]["novel_name"].(string)
	c.Data["genre9"] 	   = Rank[7]["novel_genre"].(string)
	c.Data["writer9"] 	   = Rank[7]["novel_writer"].(string)

	c.Data["id10"] 		   = Rank[8]["id"].(string)
	c.Data["name10"]  	   = Rank[8]["novel_name"].(string)
	c.Data["genre10"] 	   = Rank[8]["novel_genre"].(string)
	c.Data["writer10"] 	   = Rank[8]["novel_writer"].(string)

	c.Data["id11"] 		   = Rank[9]["id"].(string)
	c.Data["name11"]  	   = Rank[9]["novel_name"].(string)
	c.Data["genre11"] 	   = Rank[9]["novel_genre"].(string)
	c.Data["writer11"] 	   = Rank[9]["novel_writer"].(string)

	c.Data["id12"] 		   = Rank[10]["id"].(string)
	c.Data["name12"]  	   = Rank[10]["novel_name"].(string)
	c.Data["genre12"] 	   = Rank[10]["novel_genre"].(string)
	c.Data["writer12"] 	   = Rank[10]["novel_writer"].(string)

	c.Data["id13"] 		   = Rank[11]["id"].(string)
	c.Data["name13"]  	   = Rank[11]["novel_name"].(string)
	c.Data["genre13"] 	   = Rank[11]["novel_genre"].(string)
	c.Data["writer13"] 	   = Rank[11]["novel_writer"].(string)

	c.Data["id14"] 		   = Rank[12]["id"].(string)
	c.Data["name14"]  	   = Rank[12]["novel_name"].(string)
	c.Data["genre14"] 	   = Rank[12]["novel_genre"].(string)
	c.Data["writer14"] 	   = Rank[12]["novel_writer"].(string)


	c.Data["Username"] = c.Loginuser
	c.TplName = "home.html"
}
