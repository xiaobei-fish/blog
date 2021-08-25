package controllers

import (
	"fmt"
	"github.com/KenmyZhang/aliyun-communicate"
	"github.com/astaxie/beego"
)

type ForgetController struct {
	JudgeController
}

func (c *ForgetController) Get() {
	c.TplName = "forget.html"

}

func (c *ForgetController) Post() {
	c.Controller.EnableRender = false

	var (
		gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
		accessKeyId     = "Xw"
		accessKeySecret = "fY9T9TmHeOuP"
		phoneNumbers    = "15060301341"
		signName        = "blog"
		templateCode    = "SMS_149101793"
		templateParam   = "{\"code\":\"wilson\"}"
	)
	fmt.Println("phone:" + c.GetString("phone"))

	smsClient := aliyunsmsclient.New(gatewayUrl)
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	fmt.Println("Got raw response from server:", string(result.RawResponse))
	if err != nil {
		beego.Info("配置有问题")
	}

	if result.IsSuccessful() {
		fmt.Println("短信已经发送")
	} else {
		fmt.Println("短信发送失败")
	}

}
