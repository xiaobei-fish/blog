package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"web/utils"
)

func ShowPictureOnHomeByAdmin(id string) []orm.Params{
	sql := "select novel_pic from novel_info where id="
	find := "'" + id + "'"
	sql = sql + find

	var Picture []orm.Params

	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Picture)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}
	fmt.Println("Picture:",Picture)

	return Picture
}

//更新小说
func UpdateNovel(id, name, state, genre, introduction string) (int64, error) {
	orm.Debug = true
	sql := "update novel_info SET id=" + id + " , "
	name = "novel_name='" + name + "' , "
	state = "novel_state='" + state + "' , "
	genre = "novel_genre='" + genre + "' , "
	introduction = "novel_introduction='" + introduction + "' "
	location := "where id=" + id
	sql = sql + name + state + genre + introduction + location
	fmt.Println("sql::::",sql)
	return utils.ModifyDB(sql)
}

//删除文章
func DeleteNovel(novelID int) (int64, error) {
	i, err := DeleteNovelWithId(novelID)

	//计算小说总页码数,看看会不会删除导致页码减少
	SetNovelRowsNum()
	return i, err
}

func DeleteNovelWithId(ID int) (int64, error) {
	return utils.ModifyDB("delete from novel_info where id=?", ID)
}

//删除用户
func DeleteUser(userID int) (int64, error) {
	i, err := DeleteUserWithId(userID)

	//计算用户总页码数,看看会不会删除导致页码减少
	SetUserRowsNum()
	return i, err
}

func DeleteUserWithId(ID int) (int64, error) {
	return utils.ModifyDB("delete from users where id=?", ID)
}