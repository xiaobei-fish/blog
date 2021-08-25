package controllers

import (
	"blog/models"
	"fmt"
)

type SearchController struct {
	JudgeController
}

func (c *SearchController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}

	key := c.GetString("key")
	fmt.Println("key:", key)
	page, _ := c.GetInt("page")
	fmt.Println("page:",page)
	var blogList []models.Blog

	if page <= 0 {
		page = 1
	}
	//设置分页
	blogList, _ = models.FindBlogWithPage(page, key)
	fmt.Println(blogList)
	c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page, key)
	c.Data["HasFooter"] = true

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Key"] = key
	c.Data["Content"] = models.MakeSearchBlocks(blogList)

	c.TplName = "search.html"
}




//弃置方案
/*
func (c *SearchController) Get(){
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}

	key  := c.GetString("key")
	page, _ := c.GetInt("page")
	fmt.Println("key:",key)
	fmt.Println("page:",page)
	c.Data["Key"] = key

	if page == 1{
		c.Data["HasLeft"] = false
		c.Data["LeftPage"] = page - 1
		c.Data["RightPage"]= page + 1
	}else{
		c.Data["HasLeft"] = true
				c.Data["LeftPage"] = page - 1
		c.Data["RightPage"]= page + 1
	}
	//这部分爬虫可以换Python跑了存TXT再用Go读？ Go爬虫老被杀(Python不容易死？)

	//*********************爬虫**************************
	url := fmt.Sprintf("https://www.baidu.com/s?wd=%s&pn=%d",key,(page-1)*10)
	fmt.Println("url:",url)
	client := &http.Client{}
	//cookie加入,避免反爬虫
		jar, _ := cookiejar.New(nil)
	client.Jar = jar
	req, _ := http.NewRequest("GET", url, nil)
	useragent := getUserAgent()
	fmt.Printf("user-anget:%s\n",useragent)
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "www.baidu.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	res, _ := client.Do(req)
	body,_ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	htmlstr := string(body)
	//fmt.Printf("%s",body)
	reg := regexp.MustCompile(`com/link(.*?)".*?target="_blank"`)
	sul := reg.FindAllStringSubmatch(htmlstr, -1)
	//***************************************************

	var URL map[string] string
	URL = make(map[string] string)

	fmt.Printf("link有%d\n",len(sul))
	if len(sul) == 0 {
		c.Data["Content"] = models.MakeSearchBlocks(URL, false, c.IsLogin, page)
	}else {
		for i := 0; i < len(sul); i++ {
			fmt.Printf("%d,%s\n", len(sul), sul[i][1])
			URL[getTitle("https://www.baidu.com/link" + sul[i][1])] = "https://www.baidu.com/link" + sul[i][1]
		}
		for title := range URL {
			fmt.Println(title, "URL是", URL[title])
		}
		//根据得到的数据排版
		c.Data["Content"] = models.MakeSearchBlocks(URL, true, c.IsLogin, page)
	}

	c.TplName = "search.html"
}

func getUserAgent() string{
	var USER_AGENTS = [...]string {
		"'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)'",
			"'Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)'",
			"'Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)'",
			"'Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)'",
			"'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)'",
			"'Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)'",
			"'Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)'",
			"'Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)'",
			"'Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6'",
			"'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1'",
			"'Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0'",
			"'Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5'",
			"'Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6'",
			"'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11'",
			"'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20'",
			"'Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52'",
			"'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36'",
	}
	return  USER_AGENTS[rand.Intn(len(USER_AGENTS))]
}

func getTitle(url string) string{
	client := &http.Client{}
	//cookie加入,避免反爬虫
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	req, _ := http.NewRequest("GET", url, nil)
	useragent := getUserAgent()
	fmt.Printf("user-anget:%s\n",useragent)
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "www.baidu.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*//*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	res, _ := client.Do(req)
	body,_ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	htmlstr := string(body)
	//fmt.Printf("htmlstr:%s\n",htmlstr)

	//这个转码，要知道页面的编码爬虫有点难搞
	reg0 := regexp.MustCompile(`charset=([\s\S]*?)".*?>[\s\S]*?<meta http-equiv="X-UA-Compatible"`)
	char := reg0.FindAllStringSubmatch(htmlstr, -1)
	fmt.Println("charset=",char[0][1])
	htmlstr = ConvertToString(htmlstr,char[0][1],"utf-8")

	reg1 := regexp.MustCompile(`<title>([\s\S]*?)</title>`)
	sul := reg1.FindAllStringSubmatch(htmlstr, -1)
	fmt.Println(sul[0][1])

	return sul[0][1]
}
//转码，但是得知道页面编码，每个页面charset格式不一样，不好爬
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
*/