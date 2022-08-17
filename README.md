# blog

使用**Beego框架**进行一个简易blog网站的编程，包括增删改查博客以及社区和模糊查询搜索的实现。 

可优化：

通过redis中间件缓存热点key，也可以通过redis中间件维护首页热门博客排行榜。

模糊搜索查询算法优化。



#### app.conf配置

```shell
appname = web_collect
httpport = 端口
runmode = dev

#mysql配置
driverName = mysql
mysqluser = 用户名
mysqlpwd = 密码
host = 127.0.0.1
port = 3306
dbname = 数据库

#session
Sessionon = true
sessionprovider = "file"
sessionname = "web"
sessiongcmaxlifetime = 5400
sessionproviderconfig = "./tmp"
sessioncookielifetime = 5400

#翻页配置,一页5个数据
blogListPageNum = 5

```

