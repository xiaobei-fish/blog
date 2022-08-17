package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"web/utils"
)


type Collect struct {
	UserId string
	BookId string
}

//用户收藏小说
func CollectNovel(novel Collect) (int64, error) {
	i, err := insertNovel(novel)
	SetNovelRowsNum()
	return i, err
}

//储存收藏的小说id和用户名
func insertNovel(novel Collect) (int64, error) {
	return utils.ModifyDB("insert into user_collect(novel_ids,user_ids) values(?,?)",
		novel.BookId,novel.UserId)
}

//根据用户的收藏来更新rank的值以便于首页的展示
func UpdateRankByCollectNum(id string) (int64, error) {
	sql := "update novel_info set novel_rank=novel_rank + '1' where id=" + "'" + id + "'"
	fmt.Println("update:",sql)
	return utils.ModifyDB(sql)
}

//根据用户的收藏来更新rank的值以便于首页的展示,这个是用户删除的model代码
func UpdateRankByDelete(id string) (int64, error) {
	sql := "update novel_info set novel_rank=novel_rank - '1' where id=" + "'" + id + "'"
	fmt.Println("update:",sql)
	return utils.ModifyDB(sql)
}

//直接获取排名第一的数据
func GetNo1Novel() []orm.Params {
	orm.Debug = true
	var No1 []orm.Params
	sql := "select id,novel_pic,novel_name,novel_introduction,novel_writer,novel_genre from novel_info order by novel_rank desc limit 1"
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&No1)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}

	return No1
}

//根据收藏量来发送数据给首页的排行榜
func GetRankNovels() []orm.Params {
	orm.Debug = true
	var Rank []orm.Params
	sql := "select id,novel_name,novel_writer,novel_genre from novel_info order by novel_rank desc limit 1,13"
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Rank)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}

	return Rank
}

//根据Id来返回书籍的数据,用于详细页
func FindNovelWithId(Id string) []orm.Params {
	orm.Debug = true

	sql := "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info where id=" + "'" + Id + "'"

	var Novel []orm.Params

	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	fmt.Println("var是正常的")

	return Novel
}

//根据页码查询小说并返回数据以便展示
func FindNovelWithPage(page int,writer string) ([]orm.Params,int64) {
	orm.Debug = true
	//从配置文件中获取每页展示的小说数量
	num, _ := beego.AppConfig.Int("novelListPageNum")
	page--

	location := "where novel_writer=" + "'" + writer + "' "
	//fmt.Println("page::::::::::", page)
	//fmt.Println("num:::::::::::", num)
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	//fmt.Println("sql:::::::::::",sql)
	sql = "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info " + location + sql
	//fmt.Println("sql:::::::::::",sql)

	var Novel []orm.Params
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	//fmt.Println("Novel:",Novel)

	return Novel, i
}

//管理员页面
func FindNovelWithPageLimit(page, limit int) []orm.Params {
	orm.Debug = true
	page--
	//fmt.Println("page::::::::::", page)
	//fmt.Println("num:::::::::::", num)
	sql := fmt.Sprintf("limit %d,%d", page*limit, limit)
	//fmt.Println("sql:::::::::::",sql)
	sql = "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info " + sql
	//fmt.Println("sql:::::::::::",sql)

	var Novel []orm.Params
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	//fmt.Println("Novel:",Novel)

	return Novel
}

//空关键字返回全部书籍
func FindNovelWithKeyWordsOnEmpty(page int) ([]orm.Params,int64) {
	orm.Debug = true
	num, _ := beego.AppConfig.Int("novelListPageNum")
	page--
	Limit := fmt.Sprintf("limit %d,%d", page*num, num)

	var Novel []orm.Params

	sql := "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info "
	count := sql //统计匹配的总数
	sql = sql + Limit

	I, E := db.Raw(count).Values(&Novel)
	fmt.Println("输出是:",I)
	if E != nil {
		fmt.Println("Raw出错")
	}

	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	//fmt.Println("Novel:",Novel)

	return Novel, I
}

//根据关键字搜索小说
func FindNovelWithKeyWords(keyWords string, page int) ([]orm.Params,int64) {
	orm.Debug = true
	println("time:", time.Now().UnixNano())
	var Novel []orm.Params
	num, _ := beego.AppConfig.Int("novelListPageNum")
	page--
	Limit := fmt.Sprintf("limit %d,%d", page*num, num)

	sql := "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info where novel_name like "
	keyWords = "'%" + keyWords + "%' "
	sql = sql + keyWords
	writer := "or novel_genre like " + keyWords
	state := "or novel_state like " + keyWords
	introduction := "or novel_genre like " + keyWords
	sql = sql + writer + state + introduction
	count := sql //统计匹配的总数
	sql = sql + Limit//匹配搜索
	//fmt.Println("sql:::::",sql)
	I, E := db.Raw(count).Values(&Novel)
	fmt.Println("输出是:",I)
	if E != nil {
		fmt.Println("Raw出错")
	}
	println("time:", time.Now().UnixNano())
	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	//fmt.Println("Novel:",Novel)
	println("time:", time.Now().UnixNano())
	return Novel, I
}

//得到用户收藏小说的数量
func GetCollectedNumByUserId(userId string) int64  {
	cout := "select id from user_collect where user_ids="
	cout  = cout + userId

	var Novel []orm.Params

	I, E := db.Raw(cout).Values(&Novel)
	fmt.Println("输出是:",I)
	if E != nil {
		fmt.Println("Raw出错")
	}
	return I
}

//得到用作家创作小说的数量
func GetCreatedNovelNumByWriterName(writerName string) int64  {
	cout := "select id from novel_info where novel_writer=" + "'" + writerName + "'"

	var Novel []orm.Params

	I, E := db.Raw(cout).Values(&Novel)
	fmt.Println("输出是:",I)
	if E != nil {
		fmt.Println("Raw出错")
	}
	return I
}

//根据收藏的Id来返回收藏书籍的数据
func FindNovelWithCollectId(page int, userId string) ([]orm.Params,int64) {
	orm.Debug = true
	//从配置文件中获取每页展示的小说数量
	num, _ := beego.AppConfig.Int("novelListPageNum")
	page--
	//fmt.Println("page::::::::::", page)
	//fmt.Println("num:::::::::::", num)
	condition := "(select novel_ids from user_collect where user_ids="
	userid := userId + ")"
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	//fmt.Println("sql:::::::::::",sql)
	sql = "select id,novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_pic from novel_info where id in " + condition + userid + sql

	//返回用户一共收藏的数量
	I := GetCollectedNumByUserId(userId)

	var Novel []orm.Params

	i, e := db.Raw(sql).Values(&Novel)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	fmt.Println("var是正常的")

	return Novel,I
}

//存储表的行数，只有自己可以更改，当收藏夹中小说新增或者删除时需要更新这个值
var novelcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetNovelRowsNum() int {
	if novelcileRowsNum == 0 {
		novelcileRowsNum = QueryNovelRowNum()
	}
	return novelcileRowsNum
}

//查询文章的总条数
func QueryNovelRowNum() int {
	row := utils.QueryRowDB("select count(id) from novel_info")
	num := 0
	row.Scan(&num)
	return num
}

//设置小说页数
func SetNovelRowsNum(){
	novelcileRowsNum = QueryNovelRowNum()
}

//插入新小说
func AddNovel(novel NovelInfo) (int64, error) {
	i, err := insertNewNovel(novel)
	return i, err
}
func insertNewNovel(novel NovelInfo) (int64, error) {
	return utils.ModifyDB("insert into novel_info(novel_name,novel_genre,novel_introduction,novel_state,novel_writer,novel_id,novel_pic,novel_rank) values(?,?,?,?,?,?,?,?)",
		novel.Novel_name,novel.Novel_genre,novel.Novel_introduction,novel.Novel_state,novel.Novel_writer,novel.Novel_id,novel.Novel_pic,0)
}

//删除收藏夹中的小说
func DeleteCollectedNovel(novelID string,userID string) (int64, error) {

	//删除文章
	i, err := DeleteCollectedNovelWithId(novelID,userID)

	//计算小说总页码数,看看会不会删除导致页码减少
	SetNovelRowsNum()
	return i, err
}

func DeleteCollectedNovelWithId(novelID string,userID string) (int64, error) {
	sql:= "delete from user_collect where novel_ids=" + novelID + " and user_ids=" + userID
	return utils.ModifyDB(sql)
}

//删除作家自创的小说
func DeleteCreatedNovel(novelID string) (int64, error) {

	//删除文章
	i, err := DeleteCreatedNovelWithId(novelID)

	//计算小说总页码数,看看会不会删除导致页码减少
	SetNovelRowsNum()
	return i, err
}

func DeleteCreatedNovelWithId(novelID string) (int64, error) {
	sql:= "delete from novel_info where id=" + novelID
	return utils.ModifyDB(sql)
}