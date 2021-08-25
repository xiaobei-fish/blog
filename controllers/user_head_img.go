package controllers

import (
	"blog/models"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ImgController struct {
	JudgeController
}

func (c *ImgController) Post() {
	fmt.Println("fileupload...")
	fileData, fileHeader, err := c.GetFile("upload")
	if err != nil {
		c.responseErr(err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println(fileData)
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/img")
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.responseErr(err)
		return
	}

	if fileType == "img" {
		username := c.Loginuser.(string)
		Path := "../static/img/" + fileName
		fmt.Println("Path:", Path)
		_,err := models.UpdateImg(strconv.Itoa(models.QueryUserWithUsername(username)), Path)
		if err != nil {
			c.responseErr(err)
			return
		}
	}

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	c.ServeJSON()

}

func (c *ImgController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}
