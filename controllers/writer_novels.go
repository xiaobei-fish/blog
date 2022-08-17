package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"web/models"
)

type WriterController struct {
	JudgeController
}

func (c *WriterController) Get() {
	c.Data["Username"] = c.Loginuser
	c.TplName = "writer.html"
}

func (c *WriterController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	name 		 := c.GetString("novelName")
	genre 		 := c.GetString("novelGenre")
	introduction := c.GetString("novelIntroduction")
	writer 		 := c.Loginuser.(string)
	state 		 := c.GetString("novelState")


	fileData, fileHeader, err := c.GetFile("picture")
	/*fmt.Println("name:",name)
	fmt.Println("genre:",genre)
	fmt.Println("short:",introduction)
	fmt.Println("writer:",writer)
	fmt.Println("state:",state)*/

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0} //没有得到文件的信息
		c.ServeJSON()

		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println(fileData)
	//now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileTypeFlag := 0
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中,好像差不多这些格式，多了再加就行
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" ||  fileExt == ".jpeg" {
		fileTypeFlag = 1
	}
	//文件路径
	fileDir := "static\\img"
	fileName := fileHeader.Filename
	filePathStr := filepath.Join(fileDir, fileName)
	fmt.Println("Path:",filePathStr)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0}
		c.ServeJSON()

		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0}
		c.ServeJSON()

		return
	}
	if fileTypeFlag == 1 {
		//存入数据库中
		filePathStr = "..\\" + filePathStr
		novel := models.NovelInfo{Novel_name: name, Novel_introduction: introduction, Novel_state: state, Novel_writer: writer, Novel_genre: genre, Novel_pic: filePathStr, Novel_id:strconv.Itoa(int(time.Now().Unix()))}
		_, _ = models.AddNovel(novel)
	}
	c.Data["json"] = map[string]interface{}{"code": 1}
	c.ServeJSON()
	c.Redirect("/writer",302)
}