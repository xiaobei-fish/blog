package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type NoteController struct {
	JudgeController
}

func (c *NoteController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}
	noteId := c.GetString("notedId")
	fmt.Printf("%s的留言板\n",noteId)
	c.Data["NoteName"] = models.QueryUsernameWithId(noteId)
	c.Data["NotedId"] = noteId

	noteList := models.QueryNoteWithNotedId(noteId)
	c.Data["Content"] = models.MakeNoteBlocks(noteList)

	c.TplName = "note.html"
}

func (c *NoteController) Post() {
	notedMessage := c.GetString("content")
	notedId := c.GetString("noted_id")

	fmt.Printf("mes:%s,id:%s\n",notedMessage,notedId)
	var note models.Note
	note.NotedId = notedId
	note.UserId = strconv.Itoa(models.QueryUserWithUsername(c.Loginuser.(string)))
	note.Content = notedMessage
	_,err := models.InsertNote(note)
	if err != nil {
		c.responseErr(err)
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "留言成功"}

	c.ServeJSON()
}

func (c *NoteController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}




