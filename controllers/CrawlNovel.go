package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"os"
	"path"
	"regexp"
	"strconv"
	"web/models"
)

type CrawlNovelController struct {
	beego.Controller
}

func (c *CrawlNovelController) CrawlNovel() {

	var novelInfo models.NovelInfo
	//爬虫入口Url
	number := 1
	for n := 1 ; n < 6  ; n++ {
		sUrl := "https://www.qidian.com/all?orderId=2&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0&page="
		sUrl  = sUrl + strconv.Itoa(n)  //更改page值
		for m := 0; m < 20; m++ {
			rsp := httplib.Get(sUrl)
			sNovelHtml, err := rsp.String()

			if err != nil {
				panic(err)
			}
			//小说名
			reg1 := regexp.MustCompile(`<a href="//book.qidian.com/info/.*?" target="_blank" data-eid=".*?" data-bid=".*?">([\s\S]*?)</a>`)
			novelname := reg1.FindAllStringSubmatch(sNovelHtml, -1)

			//小说作者
			reg2 := regexp.MustCompile(`<a class="name" href="//my.qidian.com/author/.*?" data-eid=".*?" target="_blank">([\s\S]*?)</a>`)
			novelwriter := reg2.FindAllStringSubmatch(sNovelHtml, -1)

			//小说类型
			reg3 := regexp.MustCompile(`<a href="//www.qidian.com/.*?" target="_blank" data-eid=".*?">(.*?)</a>`)
			novelgenre := reg3.FindAllStringSubmatch(sNovelHtml, -1)

			//小说现况
			reg4 := regexp.MustCompile(`<span >([\s\S]{1,2})</span>`)
			novelstate := reg4.FindAllStringSubmatch(sNovelHtml, -1)

			//小说封面
			reg5 := regexp.MustCompile(`<img src="(//bookcover.yuewen.com/qdbimg/349573/.*?)"></a>`)
			novelpic := reg5.FindAllStringSubmatch(sNovelHtml, -1)

			//下载封面
			imgPath := "C:\\Users\\WIN10\\go\\src\\web\\static\\img\\"
			imgUrl := "https:" + novelpic[m][1]

			filename := imgPath + path.Base(imgUrl) + strconv.Itoa(number) + ".jpg"

			req := httplib.Get(imgUrl)
			content, err := req.Bytes()

			file, err := os.Create(filename)
			file.Write(content)
			number ++

			//小说id，便于索引
			reg6 := regexp.MustCompile(`<a href="//book.qidian.com/info/(.*?)"`)
			novelid := reg6.FindAllStringSubmatch(sNovelHtml, -1)

			//小说简介
			reg7 := regexp.MustCompile(`<p class="intro">([\s\S]*?)</p>`)
			novelintroduction := reg7.FindAllStringSubmatch(sNovelHtml, -1)

			//将小说信息存入数据库
			novelInfo.Novel_name = novelname[m][1]
			novelInfo.Novel_writer = novelwriter[m][1]
			novelInfo.Novel_pic = "https:" + novelpic[m][1]
			novelInfo.Novel_state = novelstate[m][1]
			novelInfo.Novel_genre = novelgenre[m+2][1]
			novelInfo.Novel_introduction = novelintroduction[m][1]
			novelInfo.Novel_id = novelid[m][1]

			models.AddNovels(&novelInfo)
			novelInfo.Id++
		}
	}
	c.Ctx.WriteString("爬取结束...")
}
